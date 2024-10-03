package integration

import (
	"VacancyService/internal/app/redis"
	"VacancyService/internal/config"
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestRedisServiceIntegration(t *testing.T) {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Logger()

	cfg := &config.Config{
		RedisConf: config.RedisConf{
			RedisMasterAddr: "127.0.0.1:6379",
			RedisSlaveAddr:  []string{"127.0.0.1:6379"},
		},
	}

	redisService := redis.NewRedisService(log, cfg)
	assert.NotNil(t, redisService, "RedisService не должен быть nil")

	pool := redisService.GetPool()
	assert.NotNil(t, pool, "Pool не должен быть nil")

	_, err := pool.Ping(context.Background()).Result()
	assert.NoError(t, err, "Ошибка при PING запросе к Redis")
}
