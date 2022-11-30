package routes

import (
	"ecommerce/controllers"
	"ecommerce/middlewares"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is middleware it check user is login or not
// AuthorizationMiddleware('role') is middleware it check user is role to access resources

func ProductRoute(router *gin.Engine) {

	router.POST("api/v1/product/add", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.AddProduct())
	router.GET("api/v1/product", middlewares.AuthenticationMiddleware(), controllers.ViewProduct())
	router.GET("api/v1/product/:id", middlewares.AuthenticationMiddleware(), controllers.ViewProductById())
	router.PUT("api/v1/product/:id", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.UpdateProduct())
	router.DELETE("api/v1/product/:id", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.DeleteProduct())
}
