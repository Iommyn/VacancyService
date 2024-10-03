package app

import (
	"VacancyService/internal/app/postgre"
	"VacancyService/internal/app/redis"
	"VacancyService/internal/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var providers = []any{
	config.GetConfig,
	newLogger,
	newServiceDiscovery,
	postgre.NewPostgreSQLService,
	redis.NewRedisService,
	newServiceContainer,
	newGinEngine,
}

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(providers...),

		fx.Invoke(func(logger zerolog.Logger, cfg *config.Config) { go newPrometheusService(logger, cfg) }),
		fx.Invoke(func(logger zerolog.Logger, cfg *config.Config, serviceDiscovery *ServiceDiscovery) {
			serviceDiscovery.RegisterService(logger, cfg)
			go serviceDiscovery.UpdateHealthCheck(logger)
		}),
		fx.Invoke(func(logger zerolog.Logger, cfg *config.Config) { postgre.NewPostgreSQLService(logger, cfg) }),
		fx.Invoke(func(logger zerolog.Logger, cfg *config.Config) { redis.NewRedisService(logger, cfg) }),

		fx.Invoke(func(services *ServiceContainer) { newGinEngine(services) }),
	)
}
