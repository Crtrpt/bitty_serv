package api

import (
	"bitty/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SessionCreateForm struct {
	UserId   string `form:"user_id" json:"user_id" binding:"required"`
	TargetId string `form:"target_id" json:"target_id" binding:"required"`
	Type     string `json:"type" `
}

// 获取聊天的session信息
func SessionCreate(c *gin.Context) {
	fmt.Print("创建聊天室")
	var form SessionCreateForm
	if c.BindJSON(&form) == nil {
		session := new(model.Session)
		session.Name = "私人聊天室"
		session.SessionId = node.Generate().Base64()
		switch form.Type {
		case "chat":

			session.Type = 0
			engine.Insert((session))
			//创建session 成员
			var user1info = &model.User{}
			has, err := engine.Where("user_id=?", form.UserId).Get(user1info)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 1,
					"msg":  "user error",
				})
			}
			if !has {
				c.JSON(200, gin.H{
					"code": 2,
					"msg":  "user error",
				})
			}

			var user2info = &model.User{}
			has, err = engine.Where("user_id=?", form.TargetId).Get(user2info)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 1,
					"msg":  "user error",
				})
			}
			if !has {
				c.JSON(200, gin.H{
					"code": 2,
					"msg":  "user error",
				})
			}

			var user1 = &model.SessionMember{UserId: form.UserId, SessionId: session.SessionId}
			user1.SessionName = user2info.NickName
			user1.SessionAvatar = user2info.Avatar
			engine.Insert(user1)

			var user2 = &model.SessionMember{UserId: form.TargetId, SessionId: session.SessionId}
			user2.SessionName = user1info.NickName
			user2.SessionAvatar = user1info.Avatar
			engine.Insert(user2)
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": session,
			})
			return
			break
		case "group":

		}
	}
}

// 获取聊天的session信息
func SessionInfo(c *gin.Context) {
	fmt.Print("获取session信息")
	var userId = c.Request.URL.Query().Get("user_id")
	var sessionId = c.Request.URL.Query().Get("session_id")
	var sessionMember model.SessionMember = model.SessionMember{}
	var session model.Session = model.Session{}
	engine.Where("session_id= ? and user_id=? ", sessionId, userId).Get(&sessionMember)
	engine.Where("session_id= ? ", sessionId).Get(&session)
	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"user":    sessionMember,
			"session": session,
		},
	})
	return
}

// 获取聊天的session信息
func SessionList(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")

	var sessionMemberList []model.SessionMember

	err := engine.Where("user_id = ?", userId).Limit(150).Find(&sessionMemberList)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}

	var sessionIds []string

	var sessionMemberMap map[string](model.SessionMember) = make(map[string]model.SessionMember)

	for _, s := range sessionMemberList {
		sessionIds = append(sessionIds, s.SessionId)
		sessionMemberMap[s.SessionId] = s
	}

	var sessionList []model.Session

	err = engine.In("session_id", sessionIds).Find(&sessionList)

	for i, s := range sessionList {
		sessionList[i].Name = sessionMemberMap[s.SessionId].SessionName
		sessionList[i].Avatar = sessionMemberMap[s.SessionId].SessionAvatar
	}
	fmt.Print(sessionList)
	//返回session列表
	c.JSON(200, gin.H{
		"code": 0,
		"data": sessionList,
	})
	return
}
