package controllers

import (
	"ecommerce/services"

	"github.com/gin-gonic/gin"
)

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.AddProduct(c)
		return
	}
}

func ViewProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewProductById(c)
		return
	}
}

func ViewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ViewProduct(c)
		return
	}
}

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.UpdateProduct(c)
		return
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.DeleteProduct(c)
		return
	}
}
