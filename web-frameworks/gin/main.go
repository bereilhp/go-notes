package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/message", func(c *gin.Context) {
		var msg Message

		// Bind JSON body to the Message struct
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Send a response
		c.JSON(http.StatusOK, gin.H{
			"received": msg.Content,
		})
	})

	r.Run()
}
