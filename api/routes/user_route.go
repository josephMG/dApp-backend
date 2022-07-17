package routes

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/api/middlewares"
	"hardhat-backend/infrastructure"
	"hardhat-backend/lib/loggers"
)

// UserRoutes struct
type UserRoutes struct {
	logger         loggers.Logger
	handler        infrastructure.Router
	userController *controllers.UserController
	authMiddleware middlewares.JWTAuthMiddleware
	middlewares.PaginationMiddleware
}

func NewUserRoutes(
	logger loggers.Logger,
	handler infrastructure.Router,
	userController *controllers.UserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) *UserRoutes {
	return &UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}

// Setup user routes
func (s *UserRoutes) Setup() {
	s.logger.Info("Setting up routes")

	api := s.handler.Group("/api")
	api.GET("/user", s.PaginationMiddleware.Handle(), s.userController.GetUser)
	api.GET("/user/:id", s.userController.GetOneUser)
	api.POST("/user", s.userController.SaveUser)
	api.PUT("/user/:id",
		//		s.uploadMiddleware.Push(s.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
		s.userController.UpdateUser,
	)
	api.DELETE("/user/:id", s.userController.DeleteUser)

}
