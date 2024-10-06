package usecases

import (
	"VacancyService/internal/entity"
	"VacancyService/internal/repository"
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type VacancyServiceImpl struct {
	vacancyService VacancyService

	vacancyRepo repository.VacancyRepository
	logger      zerolog.Logger
}

func NewVacancyService(vacancyRepo repository.VacancyRepository, logger zerolog.Logger) VacancyService {
	return &VacancyServiceImpl{
		vacancyRepo: vacancyRepo,
		logger:      logger,
	}
}

func (svc *VacancyServiceImpl) CreateVacancy(ctx context.Context, vacancy *entity.Vacancy) error {
	return svc.vacancyRepo.CreateVacancy(ctx, vacancy)
}

func (svc *VacancyServiceImpl) GetVacancyByID(ctx context.Context, id int64) (*entity.Vacancy, error) {
	return svc.vacancyRepo.GetVacancyByID(ctx, id)
}

func (svc *VacancyServiceImpl) GetAllVacancies(ctx context.Context) ([]*entity.Vacancy, error) {
	vacancies, err := svc.vacancyRepo.GetAllVacancies(ctx)
	if err != nil {
		return nil, err
	}

	if len(vacancies) == 0 {
		return nil, fmt.Errorf("404")
	}

	return vacancies, nil
}

func (svc *VacancyServiceImpl) UpdateVacancy(ctx context.Context, vacancy *entity.Vacancy) error {
	existingVacancy, err := svc.vacancyRepo.GetVacancyByID(ctx, vacancy.ID)
	if err != nil {
		return err
	}

	existingVacancy.Title = vacancy.Title
	existingVacancy.Description = vacancy.Description

	err = svc.vacancyRepo.UpdateVacancy(ctx, existingVacancy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *VacancyServiceImpl) DeleteVacancy(ctx context.Context, id int64) error {
	return svc.vacancyRepo.DeleteVacancy(ctx, id)
}
