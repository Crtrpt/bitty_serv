package api

import (
	"bitty/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func unreadMessage(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")
	var list []model.Msg

	err := engine.Where("target_id = ? and status=?  and level > ?", userId, 0, 0).Limit(100).OrderBy("created_at desc").Find(&list)
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

	err := engine.Where("target_id = ?  ", userId).Limit(100).OrderBy("created_at desc").Find(&list)
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

type ActionForm struct {
	Id     int64 `json:"id" binding:"required"`
	Type   int   `json:"type" `
	Action int   `json:"action" `
}

const (
	AddContactType int = 0
)

const (
	CancelAction int = -1
	AgreeAction  int = 0
	RejectAction int = 1
)

//对action 进行操作
func messageAction(c *gin.Context) {
	var form ActionForm
	fmt.Print("对消息进行操作")
	err := c.BindJSON(&form)
	if err == nil {
		var msg model.Msg
		_, _ = engine.ID(form.Id).Get(&msg)
		switch msg.Type {
		case AddContactType:
			switch form.Action {
			case CancelAction:
				msg.ResultPayload = "已取消请求"
				break
			case AgreeAction:
				var sender = &model.User{}
				//获取发起人信息
				fmt.Print("同意")
				engine.Where("user_id = ?", msg.SourceId).Get(sender)
				msg.ResultPayload = "已同意添加" + sender.NickName + "为好友"

				var senderContact = &model.Contact{}
				engine.Where("user_id = ? and target_id", sender.UserId, msg.TargetId).Get(senderContact)

				contact := new(model.Contact)
				contact.Name = sender.NickName
				contact.UserId = msg.TargetId
				contact.TargetId = msg.SourceId
				//TODO session id
				contact.SessionId = senderContact.SessionId
				//TODO 给 target用户发送私人消息
				_, err = engine.Insert(contact)
				break
			case RejectAction:
				msg.ResultPayload = "已拒绝"
				fmt.Print("拒绝")
				break
			}
			break
		}
		msg.Status = 1
		msg.Result = form.Action
		engine.ID(form.Id).Update(msg)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": "",
		})
		return
	} else {
		print(err.Error())
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}

}
