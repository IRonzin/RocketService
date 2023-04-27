package handlers_v2

import (
	"fmt"

	"github.com/IRonzin/RocketService/handlers/request_models/models_v2"
	"github.com/gin-gonic/gin"
)

// These are two methods implementing manual CORS handling without CORS middleware
// https://habr.com/ru/companies/macloud/articles/553826/
// fetch(
//
//		'http://localhost:60001/v2/print',
//		{
//		  method: 'POST',
//		  headers: { 'Content-Type': 'application/json' },
//		  body: JSON.stringify({ "jobId":25001,"pages":16})
//		}
//	 ).then(resp => resp.text()).then(console.log)
func PrintHandler(c *gin.Context) {
	var pj models_v2.PrintJob
	if err := c.ShouldBindJSON(&pj); err != nil {
		c.JSON(400, "Incorrect printJob!")

		return
	}

	// implementing manual CORS handling without CORS middleware
	c.Header("Access-Control-Allow-Origin", "https://www.google.com")

	c.JSON(200, gin.H{"message": fmt.Sprintf("The PrintJob #%v is started! %v pages will be printed!", pj.JobId, pj.Pages)})
}

// implementing manual CORS handling without CORS middleware
func OptionsPrintHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-type")
	c.Header("Access-Control-Allow-Origin", "https://www.google.com")
}
