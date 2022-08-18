package api

import (
	"bitty/model"

	"github.com/gin-gonic/gin"
)

type GroupCreateForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	//描述
	Description string `form:"description" json:"description"`
	//群组头像
	Avatar string `form:"avatar" json:"avatar"`
}

//创建群组
func GroupCreate(c *gin.Context) {
	var form GroupCreateForm

	if c.BindJSON(&form) == nil {
		user_id := rdb.HGet(ctx, "token:"+c.Request.Header.Get("token"), "user_id")
		group := &model.Group{}
		group.Avatar = form.Avatar
		group.Description = form.Description
		group.Name = form.Name
		group.GroupId = node.Generate().Base64()
		group.OwnerUserId = user_id.Val()
		engine.Insert(group)

		groupmember := model.GroupMember{}
		groupmember.GroupId = group.GroupId
		groupmember.UserId = user_id.Val()
		groupmember.Type = 0
		engine.Insert(groupmember)
		c.JSON(200, gin.H{
			"code": 0,
			"data": group,
		})
		return
	}
}

func GroupList(c *gin.Context) {
	var userId = rdb.HGet(ctx, "token:"+c.Request.Header.Get("token"), "user_id")

	var groupList []model.GroupMember
	//获取我加入的群组列表
	err := engine.Where("user_id = ?", userId.Val()).Find(&groupList)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "error",
		"data": groupList,
	})
	// var userIds []string
	// var userMap map[string](map[string]any)
	// userMap = make(map[string](map[string]any))

	// for _, e := range list {
	// 	userIds = append(userIds, e.TargetId)
	// 	userMap[e.TargetId] = make(map[string]any)

	// 	userMap[e.TargetId]["contact"] = e
	// }

	// var userList []model.User

	// err = engine.In("user_id", userIds).Find(&userList)
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": 1,
	// 		"msg":  "userinfo error",
	// 	})
	// 	return
	// }

	// for _, e := range userList {
	// 	userMap[e.UserId]["user"] = e
	// }
	// c.JSON(200, gin.H{
	// 	"code": 0,
	// 	"data": userMap,
	// })
}
