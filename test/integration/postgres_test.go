package integration

import (
	"VacancyService/internal/app/postgre"
	"VacancyService/internal/config"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestPostgreSQLServiceIntegration(t *testing.T) {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Logger()

	cfg := &config.Config{
		PostgreConf: config.PostgreConf{
			PostgresMaster: "postgres://postgres:12345@127.0.0.1:5432/gamesparks_db?sslmode=disable",
			PostgresReplicas: []string{
				"postgres://postgres:12345@127.0.0.1:5432/gamesparks_db?sslmode=disable",
			},
		},
	}

	postgresService := postgre.NewPostgreSQLService(log, cfg)
	assert.NotNil(t, postgresService, "PostgreSQLService не должен быть nil")

	masterDB := postgresService.GetMaster()
	assert.NotNil(t, masterDB, "Мастер подключение не должно быть nil")

	err := masterDB.Ping()
	if err != nil {
		t.Fatalf("Ошибка при проверке мастер подключения: %v", err)
	}

	replicaDB := postgresService.GetReplica()
	if replicaDB != nil {
		err = replicaDB.Ping()
		if err != nil {
			t.Fatalf("Ошибка при проверке реплика подключения: %v", err)
		}
	} else {
		t.Fatal("Реплика подключение не должно быть nil")
	}

	err = masterDB.Close()
	assert.NoError(t, err, "Ошибка при закрытии мастер подключения")

	if replicaDB != nil {
		err = replicaDB.Close()
		assert.NoError(t, err, "Ошибка при закрытии реплика подключения")
	}
}
