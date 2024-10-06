package repository

import (
	"VacancyService/internal/entity"
	"context"
)

type (
	VacancyRepository interface {
		CreateVacancy(ctx context.Context, vacancy *entity.Vacancy) error

		GetVacancyByID(ctx context.Context, id int64) (*entity.Vacancy, error)
		GetAllVacancies(ctx context.Context) ([]*entity.Vacancy, error)

		UpdateVacancy(ctx context.Context, vacancy *entity.Vacancy) error

		DeleteVacancy(ctx context.Context, id int64) error
	}
)
