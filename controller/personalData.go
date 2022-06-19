package controller

import (
	"errors"

	"github.com/Favoree-Team/server-non-user-api/entity"
	"github.com/Favoree-Team/server-non-user-api/service"
	"github.com/Favoree-Team/server-non-user-api/utils"
	"github.com/gin-gonic/gin"
)

type personalDataController struct {
	personalDataService service.PersonalDataService
}

func NewPersonalDataController(personalDataService service.PersonalDataService) *personalDataController {
	return &personalDataController{
		personalDataService: personalDataService,
	}
}

// v1.GET("/personal_data/validate")
func (pc *personalDataController) ValidateJWTPersonalData(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(200, entity.ValidationResponse{
			Validation:    entity.ValidationInvalid,
			MessageDetail: "Id not found",
		})
		return
	}

	ipAddress, ok := c.Get("ip_address")
	if !ok {
		c.JSON(200, entity.ValidationResponse{
			Validation:    entity.ValidationInvalid,
			MessageDetail: "ip address not found",
		})
		return
	}

	validationResponse, err := pc.personalDataService.ValidationPersonalData(id.(string), ipAddress.(string))
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, validationResponse)
}

// v1.POST("/personal_data")
func (pc *personalDataController) CreatePersonalData(c *gin.Context) {
	var data entity.CreatePersonalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	personalDataResponse, err := pc.personalDataService.CreatePersonalData(data)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, personalDataResponse)
}

// v1.POST("/personal_data/subscribe")
func (pc *personalDataController) SubscribePersonalData(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("id not found")))
		return
	}

	ipAddress, ok := c.Get("ip_address")
	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("ip_address not found")))
		return
	}

	token, _ := c.Get("VALID_TOKEN")

	var data entity.CreateDetailPersonalData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	personalDataResponse, err := pc.personalDataService.SubscribePersonalData(id.(string), ipAddress.(string), token.(string), data)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, personalDataResponse)
}
