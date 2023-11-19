package routes

import (
	"ecommerce/controllers"
	"ecommerce/middlewares"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is middleware it check user is login or not
// AuthorizationMiddleware('role') is middleware it check user is role to access resources

func OrderRoute(router *gin.Engine) {

	router.POST("api/v1/order/add", middlewares.AuthenticationMiddleware(), controllers.AddOrder())
	router.GET("api/v1/order", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.ViewOrder())
	router.GET("api/v1/order/:id", middlewares.AuthenticationMiddleware(), controllers.ViewOrderById())
	router.PUT("api/v1/order/:id", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.UpdateOrder())
	router.DELETE("api/v1/order/:id", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.DeleteOrder())
}
