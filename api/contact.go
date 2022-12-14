package api

import (
	"bitty/model"

	"github.com/gin-gonic/gin"
)

func list(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")

	var list []model.Contact

	err := engine.Where("user_id = ?", userId).Find(&list)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}

	var userIds []string
	var userMap map[string](map[string]any)
	userMap = make(map[string](map[string]any))

	for _, e := range list {
		userIds = append(userIds, e.TargetId)
		userMap[e.TargetId] = make(map[string]any)

		userMap[e.TargetId]["contact"] = e
	}

	var userList []model.User

	err = engine.In("user_id", userIds).Find(&userList)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "userinfo error",
		})
		return
	}

	for _, e := range userList {
		userMap[e.UserId]["user"] = e
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": userMap,
	})
}

type removeContactForm struct {
	UserId   string `form:"user_id" json:"user_id" binding:"required"`
	TargetId string `form:"target_id" json:"target_id" binding:"required"`
}

func removeContact(c *gin.Context) {
	var form removeContactForm
	if c.BindJSON(&form) == nil {
		var contacttest = &model.Contact{}
		engine.Where("user_id = ? and target_id=?", form.UserId, form.TargetId).Delete(contacttest)
		c.JSON(200, gin.H{
			"code": 0,
			"data": "",
		})
		return
	}

}

type AddContactForm struct {
	UserId   string `form:"user_id" json:"user_id" binding:"required"`
	TargetId string `form:"target_id" json:"target_id" binding:"required"`
}

func addContact(c *gin.Context) {

	var form AddContactForm
	if c.BindJSON(&form) == nil {
		var userId = form.UserId

		var contacttest = &model.Contact{}

		has, err := engine.Where("user_id = ? and target_id=?", userId, form.TargetId).Get(contacttest)
		if has {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "contatc already exists ",
			})
			return
		}

		var user = &model.User{}

		has0, err := engine.Where("user_id = ?", userId).Get(user)
		if !has0 || err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "target user not found",
			})
			return
		}

		var targetId = form.TargetId

		var target = &model.User{}

		println("%s %s", userId, targetId)
		has1, err := engine.Where("user_id = ?", targetId).Get(target)
		if !has1 || err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "target user not found",
			})
			return
		}
		//????????????????????????
		contact := new(model.Contact)
		contact.Name = target.NickName
		contact.UserId = user.UserId
		contact.TargetId = target.UserId
		contact.SessionId = node.Generate().Base64()

		//TODO ??? target????????????????????????
		_, err = engine.Insert(contact)

		msg := new(model.Msg)
		msg.Content = user.NickName + "????????????????????????"
		msg.SourceId = contact.UserId
		msg.TargetId = contact.TargetId
		msg.Status = 0
		//????????????
		msg.Type = 0
		msg.Level = 1

		//TODO ??? target????????????????????????
		_, err = engine.Insert(msg)

		msg1 := new(model.Msg)
		msg1.Content = "??????" + target.NickName + "?????????????????????"
		msg1.SourceId = contact.UserId
		msg1.TargetId = contact.UserId
		msg.Status = 1
		//????????????
		msg.Type = 1
		msg.Level = 0

		//TODO ??? source?????????
		_, err = engine.Insert(msg1)

		c.JSON(200, gin.H{
			"code": 0,
			"data": "",
		})
		return

	}

}

func infoContact(c *gin.Context) {

	var userId = c.Request.URL.Query().Get("user_id")
	var targetId = c.Request.URL.Query().Get("target_id")
	var contact = &model.Contact{}
	has, _ := engine.Where("user_id = ? and target_id = ?", userId, targetId).Get(contact)
	if has {
		c.JSON(200, gin.H{
			"code": 0,
			"data": contact,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "not found record",
			"data": contact,
		})
	}

	return
}
