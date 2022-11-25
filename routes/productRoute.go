package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {

	router.POST("api/v1/product/add", controllers.AddProduct())
	router.GET("api/v1/product", controllers.ViewProduct())
	router.GET("api/v1/product/:id", controllers.ViewProductById())
	router.PUT("api/v1/product/:id", controllers.UpdateProduct())
	router.DELETE("api/v1/product/:id", controllers.DeleteProduct())
}
