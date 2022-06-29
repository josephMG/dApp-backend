package routes

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/api/middlewares"
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"
)

// UserRoutes struct
type UserRoutes struct {
	logger         loggers.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/user", s.userController.GetUser)
		api.GET("/user/:id", s.userController.GetOneUser)
		api.POST("/user", s.userController.SaveUser)
		api.POST("/user/:id", s.userController.UpdateUser)
		api.DELETE("/user/:id", s.userController.DeleteUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger loggers.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}
