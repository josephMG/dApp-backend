package routes

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/infrastructure"
	"hardhat-backend/lib/loggers"
)

// ContractRoutes struct
type ContractRoutes struct {
	logger             loggers.Logger
	handler            infrastructure.Router
	contractController controllers.ContractController
}

// Setup contract routes
func (s *ContractRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Group("/api")
	{
		api.GET("/contract/balance/:address", s.contractController.GetBalance)
		api.GET("/contract/greeting", s.contractController.GetGreeting)
		api.POST("/contract/greeting", s.contractController.PostGreeting)
	}
}

// NewContractRoutes creates new contract controller
func NewContractRoutes(
	logger loggers.Logger,
	handler infrastructure.Router,
	contractController controllers.ContractController,
) *ContractRoutes {
	return &ContractRoutes{
		handler:            handler,
		logger:             logger,
		contractController: contractController,
	}
}
