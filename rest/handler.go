package rest

import (
	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

func handlePingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
