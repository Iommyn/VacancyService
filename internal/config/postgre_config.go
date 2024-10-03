package config

type PostgreConf struct {
	PostgresMaster   string   `env:"POSTGRES_MASTER"`
	PostgresReplicas []string `env:"POSTGRES_REPLICAS" envSeparator:","`
}
