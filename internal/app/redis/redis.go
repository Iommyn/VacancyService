package redis

import (
	"VacancyService/internal/config"
	pool "github.com/bitleak/go-redis-pool/v3"
	"github.com/rs/zerolog"
)

type RedisService struct {
	pool *pool.Pool
}

func NewRedisService(log zerolog.Logger, cfg *config.Config) *RedisService {
	pool, err := pool.NewHA(&pool.HAConfig{
		Master:           cfg.RedisConf.RedisMasterAddr,
		Slaves:           cfg.RedisConf.RedisSlaveAddr,
		Password:         "",
		ReadonlyPassword: "",
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Redis: Failed to create pool")
		return nil
	}

	log.Info().Msg("Redis: Pool created successfully")

	return &RedisService{
		pool: pool,
	}
}

func (r *RedisService) GetPool() *pool.Pool {
	return r.pool
}
