package controllers

import (
	"hardhat-backend/api_errors"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/services"
	"hardhat-backend/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      loggers.Logger
	service     services.JWTAuthService
	userService *services.UserService
}

type SignInRequest struct {
	WalletAddress string `json:"walletAddress"`
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger loggers.Logger,
	service services.JWTAuthService,
	userService *services.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	jwt.logger.Info("SignIn route called")
	var request SignInRequest
	err := c.BindJSON(&request)
	if err != nil {
		utils.HandleValidationError(jwt.logger, c, api_errors.ErrInvalidRequest)
		return
	}
	walletAddress := request.WalletAddress

	user, newUser, err := jwt.userService.GetOneUserByWalletAddress(walletAddress)
	if err != nil {
		utils.HandleError(jwt.logger, c, err)
		return
	}
	// token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"data": gin.H{
			"user":    user,
			"newUser": newUser,
		},
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route",
	})
}
