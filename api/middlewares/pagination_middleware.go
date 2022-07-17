package middlewares

import (
	"hardhat-backend/constants"
	"hardhat-backend/lib/loggers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	logger loggers.Logger
}

func NewPaginationMiddleware(logger loggers.Logger) PaginationMiddleware {
	return PaginationMiddleware{logger: logger}
}

func (p PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.logger.Info("setting up pagination middleware")

		perPage, err := strconv.ParseInt(c.Query("per_page"), 10, 0)
		if err != nil {
			perPage = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil {
			page = 0
		}

		c.Set(constants.Limit, perPage)
		c.Set(constants.Page, page)

		c.Next()
	}
}
