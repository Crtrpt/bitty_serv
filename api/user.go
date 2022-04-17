package api

import (
	"github.com/gin-gonic/gin"
	"bitty/model"
)

func profile(c *gin.Context) {
	var userId =c.Request.URL.Query().Get("userId")

	var user = &model.User{UserId:userId}

	_,err := engine.Get(user)
	if(err!=nil){
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": user,
	})
}