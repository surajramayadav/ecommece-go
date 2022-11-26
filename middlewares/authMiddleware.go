package middlewares

import (
	"ecommerce/response"
	"ecommerce/security"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := strings.ReplaceAll(c.Request.Header.Get("Authorization"), "Bearer ", "")

		fmt.Println(ClientToken)
		if ClientToken == "" {
			response.SendErrorResponse(c, 401, "No Authorization Header Provided")
			c.Abort()
			return
		}
		claims, err := security.ValidateJwtToken(ClientToken)
		if err != "" {
			response.SendErrorResponse(c, 401, err)
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

func AuthorizationMiddleware(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := c.Get("role")
		if err == false {
			response.SendErrorResponse(c, 401, "No Authorization Header Provided")
			c.Abort()
			return
		}
		if role == action {
			c.Next()
		} else {
			response.SendErrorResponse(c, 401, "Not Authorized to Perform this operation")
			c.Abort()
			return
		}
	}
}
