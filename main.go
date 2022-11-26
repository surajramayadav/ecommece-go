package main

import (
	"ecommerce/config"
	"ecommerce/middlewares"
	"ecommerce/routes"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	defer log.Close()
	log.LOGGER("Test").Info("Server Started ...")
	router.NoRoute(middlewares.RouteIsNotFoundiddleware)
	router.Use(static.Serve("/images/", static.LocalFile("./uploads", false)))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Working....",
		})
	})

	router.POST("api/v1/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["images"]

		for _, file := range files {
			fmt.Println(file.Filename)

			// Retrieve file information
			extension := filepath.Ext(file.Filename)
			// Generate random file name for the new uploaded file so it doesn't override the old file with same name
			newFileName := uuid.New().String() + extension

			// Upload the file to specific dst.
			if err := c.SaveUploadedFile(file, "./uploads/"+newFileName); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	routes.UserRoute(router)
	routes.ProductRoute(router)
	routes.OrderRoute(router)

	router.Run(":" + config.PORT)
}
