package config

type PrometheusConf struct {
	PromPort int `env:"PROM_PORT" envDefault:"9090"`
}
