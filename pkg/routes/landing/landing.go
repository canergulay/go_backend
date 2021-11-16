package landing

import "github.com/gin-gonic/gin"

func MainRouter(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
