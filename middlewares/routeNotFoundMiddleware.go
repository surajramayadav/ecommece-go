package middlewares

import (
	"ecommerce/response"

	"github.com/gin-gonic/gin"
)

func RouteIsNotFoundiddleware(c *gin.Context) {
	response.SendErrorResponse(c, 404, c.Request.RequestURI+" route is not found")
	c.Abort()
	return
}
