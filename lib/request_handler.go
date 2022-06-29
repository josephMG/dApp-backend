package lib

import (
	"hardhat-backend/lib/loggers"

	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(logger loggers.Logger) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.New()
	return RequestHandler{Gin: engine}
}
