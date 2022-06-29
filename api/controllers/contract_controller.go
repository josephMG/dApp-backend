package controllers

import (
	"hardhat-backend/lib/loggers"
	"hardhat-backend/services"

	"github.com/gin-gonic/gin"
)

// ContractController data type
type ContractController struct {
	service services.ContractService
	logger  loggers.Logger
}

// NewContractController creates new user controller
func NewContractController(contractService services.ContractService, logger loggers.Logger) ContractController {
	return ContractController{
		service: contractService,
		logger:  logger,
	}
}

// GetNode gets the node
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
