package api

import "github.com/gin-gonic/gin"

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "token error",
			})
			c.AbortWithStatus(200)
			return
		}
		// rdb.Get(token)
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,Token,Platform,Version, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
