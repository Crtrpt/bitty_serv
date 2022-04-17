package api

import "github.com/gin-gonic/gin"

func list(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func search(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}
