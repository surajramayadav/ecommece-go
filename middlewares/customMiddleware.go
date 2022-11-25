package middlewares

import (
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
)

func CustomMiddleware(c *gin.Context) {
	utils.LogData("middlware called")
	c.Next()
}
