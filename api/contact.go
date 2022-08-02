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

type AddContact struct {
	UserId   string `form:"user_id" json:"user_id" binding:"required"`
	TargetId string `form:"target_id" json:"target_id" binding:"required"`
}

func add(c *gin.Context) {

	var form AddContact
	if c.BindJSON(&form) == nil {
		var userId = form.UserId

		var user = &model.User{}

		has, err := engine.Where("user_id = ?", userId).Get(user)
		if !has || err != nil {
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

		contact := new(model.Contact)
		contact.Name = target.NickName
		contact.UserId = user.UserId
		contact.TargetId = target.UserId
		//TODO 给 target用户发送私人消息
		_, err = engine.Insert(contact)

		c.JSON(200, gin.H{
			"code": 0,
			"data": "",
		})
		return

	}

}
