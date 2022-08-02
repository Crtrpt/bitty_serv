package api

import (
	"bitty/model"

	"github.com/gin-gonic/gin"
)

func unreadMessage(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")
	var list []model.Msg

	err := engine.Where("target_id = ? and status=?  and level > ?", userId, 0, 0).Find(&list)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
	})
}

func allMessage(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")
	var list []model.Msg

	err := engine.Where("target_id = ?  ", userId).Find(&list)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
	})
}
