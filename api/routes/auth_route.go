package routes

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/infrastructure"
	"hardhat-backend/lib/loggers"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger         loggers.Logger
	handler        infrastructure.Router
	authController controllers.JWTAuthController
}

// Setup user routes
func (s *AuthRoutes) Setup() {
	s.logger.Info("Setting up routes")
	auth := s.handler.Group("/api").Group("/auth")
	{
		auth.POST("/login", s.authController.SignIn)
		auth.POST("/register", s.authController.Register)
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler infrastructure.Router,
	authController controllers.JWTAuthController,
	logger loggers.Logger,
) *AuthRoutes {
	return &AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
