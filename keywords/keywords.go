package keywords

import "github.com/gin-gonic/gin"

func add(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
