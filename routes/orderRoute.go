package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {

	router.POST("api/v1/order/add", controllers.AddOrder())
	router.GET("api/v1/order", controllers.ViewOrder())
	router.GET("api/v1/order/:id", controllers.ViewOrderById())
	router.PUT("api/v1/order/:id", controllers.UpdateOrder())
	router.DELETE("api/v1/order/:id", controllers.DeleteOrder())
}
