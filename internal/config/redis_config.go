package config

type RedisConf struct {
	RedisMasterAddr string   `env:"REDIS_MASTER_ADDRESS"`
	RedisSlaveAddr  []string `env:"REDIS_SLAVE_ADDRESS" envSeparator:","`
}
