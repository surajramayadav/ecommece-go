package routes

import (
	"instant/controllers"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is middleware it check user is login or not
// AuthorizationMiddleware('role') is middleware it check user is role to access resources

func UserRoute(router *gin.Engine) {

	router.PUT("api/v1/user", controllers.UpdateUser())
	router.GET("api/v1/user/:phone_number", controllers.GetCurrentUser())
	router.GET("api/v1/user", controllers.ViewUser())
	router.POST("api/v1/user/photo/:phone_number", controllers.UploadPhoto())
}
