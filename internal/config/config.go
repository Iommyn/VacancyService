package config

import (
	"github.com/joho/godotenv"
	"go-simpler.org/env"
)

type Config struct {
	AppConf
	ConsulConf
	PrometheusConf
	RedisConf
	PostgreConf
}

func GetConfig() (*Config, error) {
	var cfg Config

	//Не забыть за коментить при пуше!
	if err := godotenv.Load("C:\\Users\\iommy\\Desktop\\VacancyService\\.env"); err != nil {
		return nil, err
	}

	if err := env.Load(&cfg, nil); err != nil {
		return nil, err
	}

	return &cfg, nil
}
