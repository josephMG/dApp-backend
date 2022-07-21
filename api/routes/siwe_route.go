package routes

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/infrastructure"
	"hardhat-backend/lib/loggers"
)

// SiweRoutes struct
type SiweRoutes struct {
	logger         loggers.Logger
	handler        infrastructure.Router
	siweController controllers.SiweController
}

// Setup siwe routes
func (s *SiweRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Group("/api")
	{
		api.POST("/siwe/verify", s.siweController.PostVerify)
	}
}

// NewSiweRoutes creates new siwe controller
func NewSiweRoutes(
	handler infrastructure.Router,
	logger loggers.Logger,
	siweController controllers.SiweController,
) *SiweRoutes {
	return &SiweRoutes{
		handler:        handler,
		logger:         logger,
		siweController: siweController,
	}
}
