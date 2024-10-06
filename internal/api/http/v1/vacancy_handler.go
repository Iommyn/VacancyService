package v1

import (
	"VacancyService/internal/api/http/v1/model"
	"VacancyService/internal/entity"
	"VacancyService/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type vacancyHandler struct {
	vacancyUsecase usecases.VacancyService
	vacancy        entity.Vacancy
	logger         zerolog.Logger
}

func NewVacancyHandler(vacancyUsecase usecases.VacancyService, logger zerolog.Logger) *vacancyHandler {
	return &vacancyHandler{vacancyUsecase: vacancyUsecase, logger: logger}
}

func (h *vacancyHandler) GetVacancyByID(ctx *gin.Context) {
	vacancyID := ctx.Param("id")
	id, err := strconv.ParseInt(vacancyID, 10, 64)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат ID", err)
		return
	}

	vacancy, err := h.vacancyUsecase.GetVacancyByID(ctx, id)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Вакансия не найдена!", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, vacancy)
}

func (h *vacancyHandler) GetAllVacancies(ctx *gin.Context) {
	vacancies, err := h.vacancyUsecase.GetAllVacancies(ctx)

	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Вакансии не найдены!", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, vacancies)
}

func (h *vacancyHandler) CreateVacancy(ctx *gin.Context) {
	var vacancy model.VacancyRequest
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		h.logger.Error().Err(err).Msg("HTTP-Handler: Ошибка десериализации структуры VacancyRequest")
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат данных", err)
		return
	}

	newVacancy := &entity.Vacancy{
		Title:       vacancy.Title,
		Description: vacancy.Description,
	}

	err := h.vacancyUsecase.CreateVacancy(ctx, newVacancy)
	if err != nil {
		switch {
		case err.Error() == "409":
			h.handleErrorResponse(ctx, http.StatusConflict, "Вакансия уже существует!", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Вакансия успешно создана"})
}

func (h *vacancyHandler) EditVacancy(ctx *gin.Context) {
	vacancyIDStr := ctx.Param("id")
	vacancyID, err := strconv.ParseInt(vacancyIDStr, 10, 64)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат ID", err)
		return
	}

	var vacancy model.VacancyRequest
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		h.logger.Error().Err(err).Msg("HTTP-Handler: Ошибка десериализации структуры VacancyRequestUpdate")
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат данных", err)
		return
	}

	updatedVacancy := &entity.Vacancy{
		ID:          vacancyID,
		Title:       vacancy.Title,
		Description: vacancy.Description,
	}

	err = h.vacancyUsecase.UpdateVacancy(ctx, updatedVacancy)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Вакансия не найдена!", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Вакансия успешно обновлена"})
}

func (h *vacancyHandler) DeleteVacancy(ctx *gin.Context) {
	vacancyIDStr := ctx.Param("id")
	vacancyID, err := strconv.ParseInt(vacancyIDStr, 10, 64)
	if err != nil {
		h.handleErrorResponse(ctx, http.StatusBadRequest, "Неверный формат ID", err)
		return
	}

	err = h.vacancyUsecase.DeleteVacancy(ctx, vacancyID)
	if err != nil {
		switch {
		case err.Error() == "404":
			h.handleErrorResponse(ctx, http.StatusNotFound, "Вакансия не найдена!", err)
		case err.Error() == "500":
			h.handleErrorResponse(ctx, http.StatusInternalServerError, "Ошибка на стороне сервера!", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Вакансия успешно удалена"})
}

func (h *vacancyHandler) handleErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	h.logger.Error().Err(err).Msgf("HTTP-Handler: %s", message)
	ctx.JSON(statusCode, model.ErrorResponse{
		Status:  statusCode,
		Message: message,
		Error:   err.Error(),
	})
}
