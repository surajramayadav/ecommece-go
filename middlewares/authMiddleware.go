package middlewares

import (
	"ecommerce/response"
	"ecommerce/security"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken != "" {
			response.SendErrorResponse(c, 401, "No Authorization Header Provided")
			c.Abort()
			return
		}
		claims, err := security.ValidateJwtToken(ClientToken)
		if err != "" {
			response.SendErrorResponse(c, 400, err)
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("id", claims.Id)
		c.Set("name", claims.Name)
		c.Set("role", claims.Role)
		c.Set("phone", claims.Phone)
		c.Next()
	}
}
