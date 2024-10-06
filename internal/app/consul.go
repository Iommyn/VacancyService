package app

import (
	"VacancyService/internal/config"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/rs/zerolog"
)

const (
	ttl     = time.Second * 8
	checkID = "check_health"
)

type ServiceDiscovery struct {
	consulClient *api.Client
}

func newServiceDiscovery(log zerolog.Logger, conf *config.Config) *ServiceDiscovery {
	log.Info().Msg("ServiceDiscovery: consul service discovery initialized")
	client, err := api.NewClient(
		&api.Config{
			Address: fmt.Sprintf("%s:%d", conf.ConsulConf.ConsulAddress, conf.ConsulConf.ConsulPort),
		},
	)
	if err != nil {
		log.Error().Err(err).Msg("ServiceDiscovery: error creating consul client")
	}
	return &ServiceDiscovery{
		consulClient: client,
	}
}

func (srv *ServiceDiscovery) UpdateHealthCheck(log zerolog.Logger) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		err := srv.consulClient.Agent().UpdateTTL(checkID, "online", api.HealthPassing)
		if err != nil {
			log.Error().Err(err).Msg("ServiceDiscovery: error updating healthcheck ttl")
		}
		<-ticker.C
	}
}

func (srv *ServiceDiscovery) RegisterService(log zerolog.Logger, conf *config.Config) {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        checkID,
	}

	register := &api.AgentServiceRegistration{
		ID:      conf.ConsulConf.ConsulId,
		Name:    conf.ConsulConf.ConsulName,
		Tags:    []string{"vacancy"},
		Address: conf.ConsulConf.ConsulServiceAddress,
		Port:    conf.ConsulConf.ConsulServicePort,
		Check:   check,
	}

	query := map[string]any{
		"type":        "service",
		"service":     conf.ConsulConf.ConsulName,
		"passingonly": true,
	}

	plane, err := watch.Parse(query)
	if err != nil {
		log.Error().Err(err).Msg("ServiceDiscovery: error parsing consul query")
	}

	plane.HybridHandler = func(index watch.BlockingParamVal, result interface{}) {
		switch msg := result.(type) {
		case []*api.ServiceEntry:
			for _, entry := range msg {
				log.Info().Any("Service", entry.Service).Msg("ServiceDiscovery: New member joined")
			}
		}
	}

	go func() {
		plane.RunWithConfig("", &api.Config{Address: fmt.Sprintf("%s:%d", conf.ConsulConf.ConsulAddress,
			conf.ConsulConf.ConsulPort)})
	}()

	err = srv.consulClient.Agent().ServiceRegister(register)

	if err != nil {
		log.Error().Err(err).Msg("ServiceDiscovery: error registering consul service")
	}

	log.Info().Msg("ServiceDiscovery: service registered")
}
