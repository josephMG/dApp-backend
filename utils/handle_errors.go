package utils

import (
	"errors"
	"hardhat-backend/api_errors"
	"hardhat-backend/lib/loggers"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func HandleValidationError(logger loggers.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func HandleErrorWithStatus(logger loggers.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

// list static errors to filter
var exceptStaticError = []error{
	gorm.ErrRecordNotFound,
	api_errors.ErrInvalidUUID,
}

// list dyanmic errors to filter
var exceptDynamicError = []error{}

// list SQL errors to filter
var exceptSQLError = []uint16{
	1062, // duplicate entry
}

var sqlError *mysql.MySQLError

func HandleError(logger loggers.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	for _, e := range exceptStaticError {
		if errors.Is(err, e) {
			return
		}
	}

	for _, e := range exceptDynamicError {
		if reflect.TypeOf(e) == reflect.TypeOf(err) {
			return
		}
	}

	if errors.As(err, &sqlError) {
		for _, code := range exceptSQLError {
			if code == sqlError.Number {
				return
			}
		}
	}

}
