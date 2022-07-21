package controllers

import (
	"hardhat-backend/api_errors"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/services"
	"hardhat-backend/utils"

	"github.com/gin-gonic/gin"
)

// SiweController data type
type SiweController struct {
	service *services.SiweService
	logger  loggers.Logger
}

type VerifyRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// NewSiweController creates new siwe controller
func NewSiweController(siweService *services.SiweService, logger loggers.Logger) SiweController {
	return SiweController{
		service: siweService,
		logger:  logger,
	}
}

// PostVerify verify message string
func (s SiweController) PostVerify(c *gin.Context) {
	var request VerifyRequest
	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleValidationError(s.logger, c, api_errors.ErrInvalidRequest)
		return
	}

	_, err = s.service.Verify(request.Message, request.Signature)
	if err != nil {
		utils.HandleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": "Verified",
	})
}
