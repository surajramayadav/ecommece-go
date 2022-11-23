package controllers

import (
	"ecommerce/response"
	"ecommerce/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserRegistartion(c)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserLogin(c)
	}
}

func ViewUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "user view by id successfully"})
	}
}

func ViewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.SendSuccessResponse(c, 200, "user data successfully")
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "user update successfully"})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "user delete successfully"})
	}
}
