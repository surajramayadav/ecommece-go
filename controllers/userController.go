package controllers

import (
	"instant/services"

	"github.com/gin-gonic/gin"
)

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserService(c)
		return
	}
}

func GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.GetUserByPhoneNumber(c)
		return
	}
}

func ViewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewAllUsers(c)
		return
	}
}

func UploadPhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserPhotoService(c)
		return
	}
}
