package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("api/v1/user/register", controllers.Register())
	router.POST("api/v1/user/login", controllers.Login())
	router.GET("api/v1/user", controllers.ViewUser())
	router.GET("api/v1/user/:id", controllers.ViewUserById())
	router.PUT("api/v1/user/:id", controllers.UpdateUser())
	router.DELETE("api/v1/user/:id", controllers.DeleteUser())
}
