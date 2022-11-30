package main

import (
	"ecommerce/config"
	"ecommerce/middlewares"
	"ecommerce/routes"
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
)

func init() {
	log.LoadConfiguration("./config/logger.json")
	fmt.Println("============================== server listening on", config.PORT, "=============================")
}

func main() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.Use(middlewares.ErrorHandler())

	router.NoRoute(middlewares.RouteIsNotFoundiddleware)

	defer log.Close()

	log.LOGGER("Test").Info("Server Started ...")

	router.Use(static.Serve("/images/", static.LocalFile("./uploads", false)))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Working....",
		})
	})

	routes.UserAuthRoute(router)
	routes.UserRoute(router)
	routes.ProductRoute(router)
	routes.OrderRoute(router)

	router.Run(":" + config.PORT)
}
