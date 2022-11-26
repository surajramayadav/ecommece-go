package routes

import (
	"ecommerce/controllers"
	"ecommerce/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("api/v1/user/register", controllers.Register())
	router.POST("api/v1/user/login", controllers.Login())
	router.GET("api/v1/user/all", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.ViewUser())
	router.GET("api/v1/user", middlewares.AuthenticationMiddleware(), controllers.ViewUserById())
	router.PUT("api/v1/user", middlewares.AuthenticationMiddleware(), controllers.UpdateUser())
	router.DELETE("api/v1/user", middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("admin"), controllers.DeleteUser())
}
