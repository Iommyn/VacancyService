package postgre

import (
	"VacancyService/internal/config"
	"database/sql"

	"github.com/rs/zerolog"
)

type PostgreSQLService struct {
	masterDB   *sql.DB
	replicasDB []*sql.DB
}

func NewPostgreSQLService(log zerolog.Logger, cfg *config.Config) *PostgreSQLService {
	service := &PostgreSQLService{}
	var err error

	log.Info().Msg("PostgreSQL: Connect to master")

	service.masterDB, err = sql.Open("postgres", cfg.PostgreConf.PostgresMaster)
	if err != nil {
		log.Fatal().Err(err).Msg("PostgreSQL: Failed to connect to master")
		return nil
	}

	log.Info().Msg("PostgreSQL: Check connection to master")
	if err := service.masterDB.Ping(); err != nil {
		log.Error().Msgf("PostgreSQL: Failed to ping to master: %v", err)
		return nil
	}

	replicaDSNs := cfg.PostgreConf.PostgresReplicas

	for _, dsn := range replicaDSNs {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Error().Err(err).Msgf("PostgreSQL: Failed to connect to replica")
			continue
		}
		if err := db.Ping(); err != nil {
			log.Error().Err(err).Msgf("PostgreSQL: Failed to connect to replica database")
			continue
		}
		service.replicasDB = append(service.replicasDB, db)
	}

	if len(service.replicasDB) == 0 {
		return nil
	}

	return service
}

func (postgre *PostgreSQLService) GetReplica() *sql.DB {
	for _, db := range postgre.replicasDB {
		if err := db.Ping(); err == nil {
			return db
		}
	}
	return nil
}

func (postgre *PostgreSQLService) GetMaster() *sql.DB {
	return postgre.masterDB
}
