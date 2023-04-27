package handlers_v2

import (
	"net/http"

	"github.com/IRonzin/RocketService/handlers/request_models/models_v2"
	"github.com/gin-gonic/gin"
)

func v2EndpointHandler(c *gin.Context) {
	c.String(200, "Обработка запроса v2 %s %s", c.Request.Method, c.Request.URL.Path)
}

func v2Add(c *gin.Context) {
	var ap models_v2.CalculatorAddParams

	if err := c.ShouldBindJSON(&ap); err != nil {
		c.JSON(403, gin.H{"error": "Calc error!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": ap.X + ap.Y})
}

func InitializeRoutes(e *gin.Engine) {
	v2 := e.Group("/v2")

	v2.GET("/products", v2EndpointHandler)
	v2.GET("/products/:productId", v2EndpointHandler)
	v2.POST("/products", v2EndpointHandler)
	v2.PUT("/products/:productId", v2EndpointHandler)
	v2.DELETE("/products/:productId", v2EndpointHandler)
	v2.POST("/add", v2Add)

	v2.POST("/print-invoice", PrintInvoiceHandler)

	//with enabled CORS
	v2.POST("/print", PrintHandler)
	//for browser CORS
	v2.OPTIONS("/print", OptionsPrintHandler)
}
