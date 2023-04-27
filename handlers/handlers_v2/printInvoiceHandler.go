package handlers_v2

import (
	"time"

	"log"
	"math/rand"

	"github.com/IRonzin/RocketService/handlers/request_models/models_v2"
	"github.com/gin-gonic/gin"
)

func PrintInvoiceHandler(c *gin.Context) {
	var pj models_v2.PrintInvoiceJob
	if err := c.ShouldBindJSON(&pj); err != nil {
		c.JSON(400, "Incorrect printInvoiceJob!")

		return
	}

	log.Printf("PrintService: creating new print job from invoice #%v...", pj.InvoiceId)
	rand.Seed(time.Now().UnixNano())
	pj.JobId = rand.Intn(1000) + 1
	log.Printf("PrintService: created print job #%v", pj.JobId)

	c.JSON(200, pj)
}
