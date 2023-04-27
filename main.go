package main

import (
	"runtime"

	"github.com/IRonzin/RocketService/handlers/handlers_v1"
	"github.com/IRonzin/RocketService/handlers/handlers_v2"
	"github.com/IRonzin/RocketService/handlers/request_models/models_v2"
	"github.com/IRonzin/RocketService/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.FindUserAgent())
	router.Use(cors.Default())

	handlers_v1.InitializeRoutes(router)
	handlers_v2.InitializeRoutes(router)

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Cors works!",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, "Операционная система %s", runtime.GOOS)
	})

	router.GET("/productJSON", func(c *gin.Context) {
		product := models_v2.Product{Id: 1, Name: "Apple"}
		c.JSON(200, product)
	})

	router.GET("/productXML", func(c *gin.Context) {
		product := models_v2.Product{Id: 2, Name: "Banana"}
		c.XML(200, product)
	})
	router.GET("/productYAML", func(c *gin.Context) {
		product := models_v2.Product{Id: 3, Name: "Mango"}
		c.YAML(200, product)
	})

	router.Run(":60001")
}
