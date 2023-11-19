package middlewares

import (
	"instant/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, ginErr := range c.Errors {
			response.SendErrorResponse(c, 500, ginErr.Error())
			return
		}
		c.Next()

	}
}
