package middlewares

import (
	"ecommerce/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			response.SendErrorResponse(c, 500, ginErr.Error())
		}
	}
}
