package routes

import (
	"instant/controllers"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is middleware it check user is login or not
// AuthorizationMiddleware('role') is middleware it check user is role to access resources

func UserAuthRoute(router *gin.Engine) {

	router.POST("api/v1/user/auth", controllers.Auth())

}
