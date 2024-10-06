package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, vacancyHandler *vacancyHandler) {
	router.Use(metricsMiddleware())

	vacancyGroupV1 := router.Group("/v1")
	{
		vacancyGroupV1.POST("/create-vacancy", vacancyHandler.CreateVacancy)

		vacancyGroupV1.GET("/vacancies", vacancyHandler.GetAllVacancies)
		vacancyGroupV1.GET("/vacancy/:id", vacancyHandler.GetVacancyByID)

		vacancyGroupV1.PATCH("/vacancy/:id", vacancyHandler.EditVacancy)

		vacancyGroupV1.DELETE("/vacancy/:id", vacancyHandler.DeleteVacancy)
	}
}
