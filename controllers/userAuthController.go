package controllers

import (
	"ecommerce/services"

	"github.com/gin-gonic/gin"
)

func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserRegistartion(c)
		return
	}
}

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserLogin(c)
		return
	}
}

func UserForgetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserForgetPassword(c)
		return
	}
}

func UserResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserResetPassword(c)
		return
	}
}

func UserLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UserLogout(c)
		return
	}
}
