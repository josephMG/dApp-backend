package controllers

import (
	"net/http"
	"strconv"

	"hardhat-backend/constants"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/models"
	"hardhat-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController data type
type UserController struct {
	service services.UserService
	logger  loggers.Logger
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger loggers.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := u.service.GetOneUser(uint(id))

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, gin.H{"data": users})
}

// SaveUser saves the user
func (u UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.WithTrx(trxHandle).CreateUser(user); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// UpdateUser updates user
func (u UserController) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"data": "user updated"})
}

// DeleteUser deletes user
func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := u.service.DeleteUser(uint(id)); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
