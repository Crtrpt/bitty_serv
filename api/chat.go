package api

import (
	"bitty/model"

	"github.com/gin-gonic/gin"
)

type MsgForm struct {
	UserId    string `form:"user_id" json:"user_id" binding:"required"`
	SessionId string `form:"session_id" json:"session_id" binding:"required"`
	Type      int    `form:"type" json:"type"`
	Content   string `form:"content" json:"content"`
	Payload   string `form:"payload" json:"payload" `
	Sn        int64  `form:"sn" json:"sn"`
}

//发送消息
func sendMsg(c *gin.Context) {
	var form MsgForm
	var err = c.BindJSON(&form)
	if err == nil {
		msg := new(model.Chat)
		msg.SessionId = form.SessionId
		msg.SenderId = form.UserId
		msg.Type = form.Type
		msg.Content = form.Content
		msg.Payload = form.Payload
		msg.Sn = int(form.Sn)
		engine.Insert(msg)
		//发送消息
		c.JSON(200, gin.H{
			"code": 0,
			"data": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  err.Error(),
	})
	return
}
