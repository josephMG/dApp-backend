package controllers

import (
	"hardhat-backend/api_errors"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/services"
	"hardhat-backend/utils"

	"github.com/gin-gonic/gin"
)

// ContractController data type
type ContractController struct {
	service services.ContractService
	logger  loggers.Logger
}

type GreetingRequest struct {
	GreetingStr string `json:"greeting"`
}

// NewContractController creates new user controller
func NewContractController(contractService services.ContractService, logger loggers.Logger) ContractController {
	return ContractController{
		service: contractService,
		logger:  logger,
	}
}

// GetBalance gets the balance from address
func (contract ContractController) GetBalance(c *gin.Context) {
	/*
		users, err := contract.service.GetAllUser()
		if err != nil {
			u.logger.Error(err)
		}
	*/
	address := c.Param("address")

	c.JSON(200, gin.H{"data": contract.service.GetBalance(address)})
}

// GetGreeting gets the greeting from contract
func (contract ContractController) GetGreeting(c *gin.Context) {
	c.JSON(200, gin.H{"data": contract.service.GetGreeting()})
}

// PostGreeting post the greeting to contract
func (contract ContractController) PostGreeting(c *gin.Context) {
	var request GreetingRequest
	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleValidationError(contract.logger, c, api_errors.ErrInvalidRequest)
		return
	}
	c.JSON(200, gin.H{"data": contract.service.PostGreeting(request.GreetingStr)})
}
