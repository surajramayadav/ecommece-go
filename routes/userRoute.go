package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is middleware it check user is login or not
// AuthorizationMiddleware('role') is middleware it check user is role to access resources

func UserRoute(router *gin.Engine) {

	router.POST("api/v1/user/add", controllers.AddUser())
	router.GET("api/v1/user/view", controllers.ViewUser())
	router.GET("api/v1/user/view/:id", controllers.ViewUserById())
	router.PUT("api/v1/user/:id", controllers.UpdateUser())
	router.DELETE("api/v1/user/:id", controllers.DeleteUser())
}
