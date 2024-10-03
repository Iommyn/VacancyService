package config

type ConsulConf struct {
	ConsulServiceAddress string `env:"CONSUL_SERVICE_ADDRESS"`
	ConsulServicePort    int    `env:"CONSUL_SERVICE_PORT"`
	ConsulAddress        string `env:"CONSUL_ADDRESS"`
	ConsulPort           uint   `env:"CONSUL_PORT"`
	ConsulName           string `env:"CONSUL_NAME"`
	ConsulId             string `env:"CONSUL_ID"`
}
