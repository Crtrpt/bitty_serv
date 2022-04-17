package api

import (
	"fmt"

	"net/http"


	"bitty/middleware"
	"bitty/model"


	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)



func Router() http.Handler {
	
	Init()
	e := gin.New()
	
	e.Use(gin.Recovery())
	e.Use(middleware.CORSMiddleware())
	v1 := e.Group("/api/v1/auth")
	{
		v1.POST("/login", login)
		v1.POST("/signup", signup)
	}
	v2 := e.Group("/api/v1/endpoint")
	{
		v2.GET("/list", list)
		v2.GET("/search", search)
	}
	v3 := e.Group("/api/v1/user")
	{
		v3.GET("/profile", profile)
	}
	return e
}

func login(c *gin.Context) {
	var form PostLogin
	if c.BindJSON(&form) == nil {
		var user = &model.User{Account: form.Account, Password: form.Password}
		var u, err = engine.Get(user)
		if err != nil {
			fmt.Printf("ERROR:%s", err)
		}
		fmt.Print(user)
		if u {
			userToken := new(model.UserToken)
			userToken.Token = node.Generate().Base64()
			userToken.UserId = user.UserId
			_, err = engine.Insert(userToken)
			if err == nil {
				c.JSON(200, gin.H{
					"code": 0,
					"data": userToken.Token,
				})
				return
			}
			fmt.Print(err)
		}

	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "user not found or password wrong",
	})

}

type PostLogin struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func signup(c *gin.Context) {
	var form PostSignup
	if c.BindJSON(&form) == nil {
		var u, err = engine.Get(&model.User{Account: form.Account})
		if err != nil {
			fmt.Printf("ERROR:%s", err)
		}
		if u {
			c.JSON(200, gin.H{
				"code": 1,
				"data": "account already exists",
			})
			return
		}
		// create account
		user := new(model.User)
		user.Account = form.Account
		user.Password = form.Password
		user.NickName = form.Account
		user.UserId = node.Generate().Base64()
		_, err = engine.Insert(user)
		if err == nil {
			userToken := new(model.UserToken)
			userToken.Token = node.Generate().Base64()
			userToken.Id = user.Id
			_, err = engine.Insert(userToken)
			if err == nil {
				c.JSON(200, gin.H{
					"code": 0,
					"data": userToken.Token,
				})
				return
			}
			fmt.Print(err)

			c.JSON(200, gin.H{
				"code": 0,
				"data": userToken.Token,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "user not found or password wrong",
	})
}

type PostSignup struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
