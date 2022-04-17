package api

import (
	"github.com/gin-gonic/gin"
	"bitty/model"
)

func profile(c *gin.Context) {
	var userId =c.Request.URL.Query().Get("userId")

	var user = &model.User{UserId:userId}

	has ,err := engine.Get(user)
	if(!has || err!=nil){
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": user,
	})
}