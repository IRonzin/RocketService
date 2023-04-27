package handlers_v1

import (
	"github.com/IRonzin/RocketService/handlers/request_models/models_v1"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(e *gin.Engine) {
	v1 := e.Group("/v1")

	v1.GET("/products", v1EndpointHandler)
	v1.GET("/products/:productId", v1EndpointHandler)
	v1.POST("/products", v1EndpointHandler)
	v1.PUT("/products/:productId", v1EndpointHandler)
	v1.DELETE("/products/:productId", v1EndpointHandler)
	v1.GET("/add/:x/:y", v1Add)
}

func v1EndpointHandler(c *gin.Context) {
	c.String(200, "Обработка запроса v1 %s %s", c.Request.Method, c.Request.URL.Path)
}

func v1Add(c *gin.Context) {
	var ap models_v1.CalculatorAddParams
	if err := c.ShouldBindUri(&ap); err != nil {
		c.JSON(403, gin.H{"error": "Calc error!"})

		return
	}

	c.String(200, "%f", ap.X+ap.Y)
}
