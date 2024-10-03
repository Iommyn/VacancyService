package repository

import (
	"VacancyService/internal/entity"
	"context"
)

type (
	VacancyRepository interface {
		GetVacancyByID(ctx context.Context, id int64) (*entity.Vacancy, error)
		CreateVacancy(ctx context.Context, vacancy *entity.Vacancy) error
		UpdateVacancy(ctx context.Context, vacancy *entity.Vacancy) error
		DeleteVacancy(ctx context.Context, id int64) error
		GetAllVacancies(ctx context.Context) ([]*entity.Vacancy, error)
	}
)
