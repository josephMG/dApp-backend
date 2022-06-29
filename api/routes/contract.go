package routes

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"
)

// ContractRoutes struct
type ContractRoutes struct {
	logger             loggers.Logger
	handler            lib.RequestHandler
	contractController controllers.ContractController
}

// Setup contract routes
func (s ContractRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/contract/get_balance/:address", s.contractController.GetBalance)
	}
}

// NewContractRoutes creates new contract controller
func NewContractRoutes(
	logger loggers.Logger,
	handler lib.RequestHandler,
	contractController controllers.ContractController,
) ContractRoutes {
	return ContractRoutes{
		handler:            handler,
		logger:             logger,
		contractController: contractController,
	}
}
