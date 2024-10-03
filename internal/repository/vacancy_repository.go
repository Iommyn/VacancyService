package repository

import (
	"VacancyService/internal/app/postgre"
	"VacancyService/internal/app/redis"
	"VacancyService/internal/entity"
	"VacancyService/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"time"
)

type vacancyRepository struct {
	db    *postgre.PostgreSQLService
	redis *redis.RedisService

	logger zerolog.Logger
}

func NewVacancyRepository(db *postgre.PostgreSQLService, redis *redis.RedisService, logger zerolog.Logger) VacancyRepository {
	return &vacancyRepository{db: db, redis: redis, logger: logger}
}

func (v vacancyRepository) UpdateVacancy(ctx context.Context, vacancy *entity.Vacancy) error {
	dbReplica := v.db.GetReplica()

	vacancyModel := &models.Vacansy{
		ID:          vacancy.ID,
		Title:       vacancy.Title,
		Description: vacancy.Description,
		UpdatedAt:   time.Now(),
	}

	rowsAffected, err := vacancyModel.Update(ctx, dbReplica, boil.Infer())
	if err != nil {
		return fmt.Errorf("500")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("404")
	}

	return nil
}

func (v vacancyRepository) GetVacancyByID(ctx context.Context, id int64) (*entity.Vacancy, error) {
	dbReplica := v.db.GetReplica()

	vacancyModel, err := models.Vacansies(qm.Where("id = ?", id)).One(ctx, dbReplica)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("vacancy not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	vacancy := &entity.Vacancy{
		ID:          vacancyModel.ID,
		Title:       vacancyModel.Title,
		Description: vacancyModel.Description,
		Updated_At:  vacancyModel.UpdatedAt.Format(time.RFC3339),
		Created_At:  vacancyModel.CreatedAt.Format(time.RFC3339),
	}

	return vacancy, nil
}

func (v vacancyRepository) DeleteVacancy(ctx context.Context, id int64) error {
	dbMaster := v.db.GetMaster()

	vacancyModel := &models.Vacansy{
		ID: id,
	}

	rowsAffected, err := vacancyModel.Delete(ctx, dbMaster)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return fmt.Errorf("409")
			}
		}
		return fmt.Errorf("500")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Вакансия не найдена: %v", http.StatusNotFound)
	}

	return nil
}

func (v vacancyRepository) CreateVacancy(ctx context.Context, vacancy *entity.Vacancy) error {
	newVacancy := models.Vacansy{
		Title:       vacancy.Title,
		Description: vacancy.Description,
	}

	dbMaster := v.db.GetMaster()

	err := newVacancy.Insert(ctx, dbMaster, boil.Infer())

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return fmt.Errorf("409")
			}
		}
		return fmt.Errorf("500")
	}

	return nil
}

func (v vacancyRepository) GetAllVacancies(ctx context.Context) ([]*entity.Vacancy, error) {
	dbReplica := v.db.GetReplica()

	vacancyModels, err := models.Vacansies().All(ctx, dbReplica)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return nil, fmt.Errorf("409")
			}
		}
		return nil, fmt.Errorf("500")
	}

	var vacancies []*entity.Vacancy
	for _, vacancyModel := range vacancyModels {
		vacancy := &entity.Vacancy{
			ID:          vacancyModel.ID,
			Title:       vacancyModel.Title,
			Description: vacancyModel.Description,
			Updated_At:  vacancyModel.UpdatedAt.Format(time.RFC3339),
			Created_At:  vacancyModel.CreatedAt.Format(time.RFC3339),
		}
		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}
