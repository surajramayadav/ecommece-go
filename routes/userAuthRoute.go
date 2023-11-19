package routes

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is middleware it check user is login or not
// AuthorizationMiddleware('role') is middleware it check user is role to access resources

func UserAuthRoute(router *gin.Engine) {

	router.POST("api/v1/user/auth/register", controllers.UserRegister())
	router.POST("api/v1/user/auth/login", controllers.UserLogin())
	router.GET("api/v1/user/auth/forget-password", controllers.UserForgetPassword())
	router.GET("api/v1/user/auth/reset-password", controllers.UserResetPassword())
	router.PUT("api/v1/user/auth/logout", controllers.UserLogout())
}
