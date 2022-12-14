package api

import (
	"bitty/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func profile(c *gin.Context) {
	var userId = c.Request.URL.Query().Get("user_id")

	var user = &model.User{UserId: userId}

	has, err := engine.Where("user_id = ? ", userId).Get(user)
	if !has || err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": user,
	})
}
func UserSession(c *gin.Context) {
	fmt.Print("获取用户的全部session")
	var userId = c.Request.URL.Query().Get("user_id")
	var sessionMember = []model.SessionMember{}
	engine.Where("user_id = ? ", userId).Find(&sessionMember)
	c.JSON(200, gin.H{
		"code": 0,
		"data": sessionMember,
	})
}
func search(c *gin.Context) {
	var keywords = c.Request.URL.Query().Get("keywords")

	var list []model.User

	var query = engine.Cols("nick_name", "user_id", "status", "avatar").Where("allow_search = ?  and nick_name like ?", 1, keywords+"%").UseBool().Limit(10)
	query.Find(&list)

	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
	})
	return
}

type SaveProfileForm struct {
	UserId           string `form:"user_id" json:"user_id" binding:"required"`
	NickName         string `form:"nick_name" json:"nick_name" binding:"required"`
	Email            string `form:"email" json:"email" binding:"required"`
	Status           string `form:"status" json:"status" binding:"required"`
	Avatar           string `form:"avatar" json:"avatar" binding:"required"`
	AllowSearch      bool   `form:"allow_search" json:"allow_search" binding:"required"`
	AllowAnonSession bool   `form:"allow_anon_session" json:"allow_anon_session" binding:"required"`
}

func save(c *gin.Context) {
	var form SaveProfileForm
	var err = c.BindJSON(&form)
	if err == nil {
		var user = &model.User{UserId: form.UserId}
		var u, err = engine.Get(user)
		if err != nil && u {
			fmt.Printf("ERROR:%s", err)
		}
		fmt.Print(user)
		user.NickName = form.NickName
		user.Status = form.Status
		user.Email = form.Email
		user.Avatar = form.Avatar
		user.AllowSearch = form.AllowSearch
		user.AllowAnonSession = form.AllowAnonSession
		engine.ID(user.Id).AllCols().Update(user)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
		return

	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  err.Error(),
	})
}
