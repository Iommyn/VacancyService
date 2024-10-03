package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, vacancyHandler *vacancyHandler) {
	router.Use(metricsMiddleware())
	vacancyGroupV1 := router.Group("/v1")
	{
		vacancyGroupV1.POST("/create-vacancy", vacancyHandler.CreateVacancy)

		vacancyGroupV1.GET("/vacancies", vacancyHandler.GetAllVacancies)
		vacancyGroupV1.GET("/vacancies/:id", vacancyHandler.GetVacancyByID)

		vacancyGroupV1.PATCH("/vacancies/:id", vacancyHandler.EditVacancy)

		vacancyGroupV1.DELETE("/vacancies/:id", vacancyHandler.DeleteVacancy)
	}
}
