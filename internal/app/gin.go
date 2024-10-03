package app

import (
	v1 "VacancyService/internal/api/http/v1"
	"VacancyService/internal/app/postgre"
	"VacancyService/internal/app/redis"
	"VacancyService/internal/config"
	"VacancyService/internal/metrics"
	"VacancyService/internal/repository"
	"VacancyService/internal/usecases"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "net/http/pprof"
)

type ServiceContainer struct {
	Config *config.Config
	Logger zerolog.Logger

	PostgreSQL *postgre.PostgreSQLService
	Redis      *redis.RedisService
}

func newServiceContainer(
	cfg *config.Config,
	logger zerolog.Logger,

	postgreSQL *postgre.PostgreSQLService,
	redis *redis.RedisService,

) *ServiceContainer {
	return &ServiceContainer{
		Config: cfg,
		Logger: logger,

		PostgreSQL: postgreSQL,
		Redis:      redis,
	}
}

func newGinEngine(services *ServiceContainer) *gin.Engine {
	log.Info().Msg("GIN-SERVER: Create gin service")

	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	srv := &http.Server{
		Addr:    services.Config.AppConf.AppListener,
		Handler: r.Handler(),
	}

	vacancyRepo := repository.NewVacancyRepository(services.PostgreSQL, services.Redis, services.Logger)
	vacancyUsecase := usecases.NewVacancyService(vacancyRepo, services.Logger)
	vacancyHandler := v1.NewVacancyHandler(vacancyUsecase, services.Logger)

	v1.RegisterRoutes(r, vacancyHandler)

	go collectMetrics()

	log.Info().Msg("GIN-SERVER: Start http server")

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("GIN-SERVER: Start http server failed")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("GIN-SERVER: Shutdown http server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("GIN-SERVER: Shutdown http server failed")
	}

	select {
	case <-ctx.Done():
		log.Info().Msg("GIN-SERVER: Shutdown http server")
	}
	log.Info().Msg("GIN-SERVER: Exit")

	return r
}

func collectMetrics() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		metrics.VacancyMemoryUsage.Set(float64(memStats.Alloc) / 1024 / 1024)
		metrics.VacancyCPUUsage.Set(float64(runtime.NumCPU()))
	}
}
