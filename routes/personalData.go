package routes

import (
	"github.com/Favoree-Team/server-non-user-api/controller"
	"github.com/Favoree-Team/server-non-user-api/repository"
	"github.com/Favoree-Team/server-non-user-api/service"

	"github.com/gin-gonic/gin"
)

var (
	PersonalDataRepository = repository.NewPersonalDataRepository(DB)
	PersonalDataService    = service.NewPersonalDataService(PersonalDataRepository, authService)
	PersonalDataController = controller.NewPersonalDataController(PersonalDataService)
)

func PersonalDataRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/personal_data/validate", MainMiddleware, PersonalDataController.ValidateJWTPersonalData)
		v1.POST("/personal_data", PersonalDataController.CreatePersonalData)
		v1.POST("/personal_data/subscribe", MainMiddleware, PersonalDataController.SubscribePersonalData)
	}
}
