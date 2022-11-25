package main

import (
	"ecommerce/config"
	"ecommerce/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	// database.Connection()

	fmt.Println("============================== server listening on", config.PORT, "=============================")
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Working....",
		})
	})
	routes.UserRoute(router)
	routes.ProductRoute(router)
	routes.OrderRoute(router)
	router.Run(":" + config.PORT)
}
