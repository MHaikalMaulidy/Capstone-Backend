package controller

import (
	"backend-summarizer/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SummarizeRequest struct {
	Text string `json:"text" binding:"required"`
	N    int    `json:"n"`
}

func HandleSummarize(c *gin.Context) {
	var req SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	n := req.N
	if n <= 0 {
		n = 2
	}

	summary := service.ProcessSummary(req.Text, n)
	c.JSON(http.StatusOK, gin.H{
		"original_text": req.Text,
		"summary_text":  summary,
		"timestamp":     c.MustGet("timestamp"),
	})
}

func GetSummaries(c *gin.Context) {
	summaries := service.FetchSummaries()
	c.JSON(http.StatusOK, gin.H{"data": summaries})
}
