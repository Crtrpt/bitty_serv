package api

import (
	"bitty/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type GroupRemoveForm struct {
	GroupId string `form:"group_id" json:"group_id" binding:"required"`
}

func GroupRemove(c *gin.Context) {
	var form GroupRemoveForm

	if c.BindJSON(&form) == nil {
		group := &model.Group{}
		engine.Where("group_id = ?", form.GroupId).Limit(1).Get(group)
		user_id := rdb.HGet(ctx, "token:"+c.Request.Header.Get("token"), "user_id")
		if group.OwnerUserId == user_id.Val() {
			//删除所有群组成员
			engine.Where("group_id = ?", form.GroupId).Delete(&model.GroupMember{})
			engine.Where("group_id = ?", form.GroupId).Delete(group)

			c.JSON(200, gin.H{
				"code": 0,
				"data": group,
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code": 1,
				"data": "你不是群管理员",
			})
			return
		}
	}
}

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
		group.SessionId = node.Generate().Base64()
		group.GroupId = node.Generate().Base64()
		group.OwnerUserId = user_id.Val()
		engine.Insert(group)

		groupmember := model.GroupMember{}
		groupmember.GroupId = group.GroupId
		groupmember.UserId = user_id.Val()
		groupmember.Type = 0
		engine.Insert(groupmember)

		//创建群组聊天室
		session := &model.Session{}
		session.Name = group.Name
		session.SessionId = group.SessionId
		session.Avatar = group.Avatar
		session.Type = 1
		session.CreatedAt = time.Now()
		_, err = engine.Insert((session))
		if err != nil {
			fmt.Printf(err.Error())
		}
		//拉取成员信息
		var userinfo = &model.User{}
		engine.Where("user_id=?", user_id.Val()).Get(userinfo)

		var user1 = &model.SessionMember{
			UserId:    user_id.Val(),
			SessionId: session.SessionId,
			Type:      0,
		}
		engine.Insert(user1)
		//创建群会话
		c.JSON(200, gin.H{
			"code": 0,
			"data": group,
		})
		return
	}
}

type GroupJoinForm struct {
	GroupId string `form:"group_id" json:"group_id" binding:"required"`
}

func GroupJoin(c *gin.Context) {
	var form GroupJoinForm
	//TODO加入群组
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "todo",
		"data": form,
	})
	return
}
func GroupList(c *gin.Context) {
	var userId = rdb.HGet(ctx, "token:"+c.Request.Header.Get("token"), "user_id")

	var groupMemberList []model.GroupMember
	//获取我加入的群组列表
	err := engine.Where("user_id = ?", userId.Val()).Find(&groupMemberList)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}

	var groupIds []string
	var groupMap map[string](map[string]any)
	groupMap = make(map[string](map[string]any))

	for _, e := range groupMemberList {
		groupIds = append(groupIds, e.GroupId)
		groupMap[e.GroupId] = make(map[string]any)
		groupMap[e.GroupId]["group"] = e
	}

	var groupList []model.Group

	err = engine.In("group_id", groupIds).Find(&groupList)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "group error",
		})
		return
	}

	for _, e := range groupList {
		groupMap[e.GroupId]["group"] = e
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": groupMap,
	})
}

func GroupInfo(c *gin.Context) {

	// var userId = c.Request.URL.Query().Get("user_id")
	// var targetId = c.Request.URL.Query().Get("target_id")
	// var contact = &model.Contact{}
	// has, _ := engine.Where("user_id = ? and target_id = ?", userId, targetId).Get(contact)
	// if has {
	// 	c.JSON(200, gin.H{
	// 		"code": 0,
	// 		"data": contact,
	// 	})
	// } else {
	// 	c.JSON(200, gin.H{
	// 		"code": 1,
	// 		"msg":  "not found record",
	// 		"data": contact,
	// 	})
	// }

	// return
}

func GroupProfile(c *gin.Context) {

	var group_id = c.Request.URL.Query().Get("group_id")

	var group = &model.Group{}
	has, _ := engine.Where("group_id = ?", group_id).Get(group)
	if has {
		c.JSON(200, gin.H{
			"code": 0,
			"data": group,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "not found record",
			"data": group,
		})
	}

	return
}

func GroupSearch(c *gin.Context) {
	var keywords = c.Request.URL.Query().Get("keywords")

	var list []model.Group

	var query = engine.Where("name like ?", keywords+"%").UseBool().Limit(10)
	query.Find(&list)

	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
	})
	return
}
