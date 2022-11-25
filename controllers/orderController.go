package controllers

import (
	"ecommerce/services"

	"github.com/gin-gonic/gin"
)

func AddOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.AddOrder(c)
		return
	}
}

func ViewOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewOrderById(c)
		return
	}
}

func ViewOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewOrder(c)
		return
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UpdateOrder(c)
		return
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.DeleteOrder(c)
		return
	}
}
