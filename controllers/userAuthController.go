package controllers

import (
	"instant/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.AuthService(c)
		return
	}
}
