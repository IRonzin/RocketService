package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func FindUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("It's logger middleware: ", c.GetHeader("User-Agent"))
		// Before calling handler
		c.Next()
		// After calling handler
	}
}
