package controllers

import (
	"hardhat-backend/api_errors"
	"hardhat-backend/lib"
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
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	paramID := c.Param("id")

	userID, err := lib.ShouldParseUUID(paramID)
	if err != nil {
		utils.HandleValidationError(jwt.logger, c, api_errors.ErrInvalidUUID)
		return
	}

	user, _ := jwt.userService.GetOneUser(userID)
	token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route",
	})
}
