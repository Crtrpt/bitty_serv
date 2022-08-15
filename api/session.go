package api

import (
	"bitty/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type SessionSuspendForm struct {
	UserId    string `form:"user_id" json:"user_id" binding:"required"`
	SessionId string `form:"session_id" json:"session_id" binding:"required"`
	Value     bool   `form:"value" json:"value"`
}

func SessionSuspend(c *gin.Context) {
	var form SessionSuspendForm
	if c.BindJSON(&form) == nil {
		var userId = form.UserId
		var sessionId = form.SessionId
		var sessionMember model.SessionMember = model.SessionMember{}
		// var session model.Session = model.Session{}
		sessionMember.Suspend = form.Value
		//退出会话
		engine.Where("session_id= ? and user_id=? ", sessionId, userId).Limit(1).Cols("suspend").Update(&sessionMember)
		c.JSON(200, gin.H{
			"code": 0,
			"data": sessionMember,
		})
		return
	}
}

type SessionRemoveForm struct {
	UserId    string `form:"user_id" json:"user_id" binding:"required"`
	SessionId string `form:"session_id" json:"session_id" binding:"required"`
}

func SessionRemove(c *gin.Context) {
	var form SessionRemoveForm
	if c.BindJSON(&form) == nil {
		var userId = form.UserId
		var sessionId = form.SessionId
		var sessionMember model.SessionMember = model.SessionMember{}
		// var session model.Session = model.Session{}
		sessionMember.DeletedAt = time.Time{}
		//退出会话
		engine.Where("session_id= ? and user_id=? ", sessionId, userId).Limit(1).Cols("deleted_at").Update(&sessionMember)
		c.JSON(200, gin.H{
			"code": 0,
			"data": "",
		})
		return
	}
}

type SessionResumeForm struct {
	UserId    string `form:"user_id" json:"user_id" binding:"required"`
	SessionId string `form:"session_id" json:"session_id" binding:"required"`
}

func SessionResume(c *gin.Context) {
	fmt.Print("恢复会话订阅")
	var form SessionResumeForm
	if c.BindJSON(&form) == nil {
		var userId = form.UserId
		var sessionId = form.SessionId
		var sessionMember model.SessionMember = model.SessionMember{}
		sessionMember.Suspend = true
		//退出会话
		engine.Where("session_id= ? and user_id=? ", sessionId, userId).Cols("suspend").Limit(1).Update(&sessionMember)
		c.JSON(200, gin.H{
			"code": 0,
			"data": sessionMember,
		})
		return
	}
}

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

		switch form.Type {
		case "chat":
			var checkSessionList = []model.SessionMember{}
			//查找我的所有聊天室
			engine.Where("user_id=?", form.UserId).Find(&checkSessionList)

			var chsids []string
			for _, e := range checkSessionList {
				chsids = append(chsids, e.SessionId)
			}

			var checkSessionMemberList = []model.SessionMember{}

			engine.In("session_id", chsids).Where("user_id = ?", form.TargetId).Find(&checkSessionMemberList)

			fmt.Print(checkSessionMemberList)
			if len(checkSessionMemberList) > 0 {

				var exitSession = model.Session{}
				engine.Where("session_id=?", checkSessionMemberList[0].SessionId).Limit(1).Get(&exitSession)
				//所有的改为未删除状态
				c.JSON(200, gin.H{
					"code": 0,
					"msg":  "ok",
					"data": exitSession,
				})
				return
			}
			//判断有没有
			session := new(model.Session)
			session.Name = "私人聊天室"
			session.SessionId = node.Generate().Base64()

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

func SessionProfile(c *gin.Context) {
	fmt.Print("获取session详细信息")
	var sessionId = c.Request.URL.Query().Get("session_id")
	var sessionMembers []model.SessionMember = []model.SessionMember{}
	engine.Cols("user_id").Where("session_id = ?", sessionId).Find(&sessionMembers)
	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"member": sessionMembers,
		},
	})
	return
}

// 获取聊天的session信息
func SessionList(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")

	var sessionMemberList []model.SessionMember

	err := engine.Where("user_id = ? ", userId).Limit(150).Find(&sessionMemberList)
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
