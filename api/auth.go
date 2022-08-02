package api

import (
	"bitty/middleware"
	"bitty/model"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
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
		v1.POST("/sendcode", sendCode)
		v1.POST("/resetpassword", resetpassword)
	}
	v2 := e.Group("/api/v1/endpoint")
	{
		v2.POST("/add", add)
		v2.GET("/list", list)
		v2.GET("/search", search)
	}
	v3 := e.Group("/api/v1/user")
	{
		v3.GET("/profile", profile)
		v3.POST("/save", save)
	}
	return e
}

func login(c *gin.Context) {
	var form PostLogin
	if c.BindJSON(&form) == nil {
		var user = &model.User{Account: form.Account}
		var u, err = engine.Get(user)
		if err != nil {
			fmt.Printf("ERROR:%s", err)
		}
		fmt.Print(user)
		serverpassword := DecryptAES([]byte(os.Getenv("encrypt_key")), user.Password)
		if serverpassword != form.Password {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "账号密码不匹配",
				"data":    "",
			})
			return
		}
		if u {
			userToken := new(model.UserToken)
			userToken.Token = node.Generate().Base64()
			userToken.UserId = user.UserId
			userToken.LastLoginIp = c.ClientIP()
			fmt.Print(c.Request.Header)
			userToken.Platform = c.Request.Header["Platform"][0]
			userToken.ClientVersion = c.Request.Header["Version"][0]
			_, err = engine.Insert(userToken)
			if err == nil {
				c.JSON(200, gin.H{
					"code": 0,
					"data": gin.H{
						"token": userToken.Token,
						"user":  user,
					},
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
				"code":    1,
				"message": "账户已经存在",
				"data":    "",
			})
			return
		}

		// create account
		user := new(model.User)
		user.Account = form.Account
		user.Password = EncryptAES([]byte(os.Getenv("encrypt_key")), form.Password)
		user.NickName = form.Account
		user.Email = form.Email
		user.UserId = node.Generate().Base64()
		fmt.Printf("创建用户")
		_, err = engine.Insert(user)
		//发送注册邮件

		if err != nil {
			fmt.Print(err)
			c.JSON(200, gin.H{
				"code":    1,
				"message": "系统异常",
				"data":    "",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "创建用户成功",
		"data":    "",
	})
	return
}

type PostSignup struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

var (
	m errgroup.Group
)

func sendCode(c *gin.Context) {
	var form SendCodeForm
	if c.BindJSON(&form) == nil {

		var u, err = engine.Get(&model.User{Account: form.Account, Email: form.Email})
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "账号不存在",
				"data":    "",
			})
		}

		fmt.Print(u)
		if u {
			fmt.Printf("发送找回邮件")
			m.Go(func() error {

				rand.Seed(time.Now().UnixNano())
				code := strconv.Itoa(rand.Intn(10000))
				err := rdb.Set(ctx, "verif:"+form.Email, code, 60*time.Second)
				if err != nil {
					fmt.Print(err)
				}
				mail_err := sendMail(form.Email, code, "重置密码验证码")
				return mail_err
			})

			c.JSON(200, gin.H{
				"code":    0,
				"message": "验证码已经发送到你的邮箱",
				"data":    "",
			})
			return
		}
		//发送注册邮件
	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "邮箱不存在",
	})
}

type SendCodeForm struct {
	Account string `form:"account" json:"account" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required"`
}

func resetpassword(c *gin.Context) {
	var form resetpasswordForm
	if c.BindJSON(&form) == nil {
		var user = model.User{Account: form.Account, Email: form.Email}
		var u, err = engine.Get(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "账号不存在",
				"data":    "",
			})
		}
		if u {

			code, err := rdb.Get(ctx, "verif:"+form.Email).Result()
			if err != nil || code != form.Code {
				c.JSON(200, gin.H{
					"code":    1,
					"message": "验证码错误或者已过期",
					"data":    "",
				})
				return
			}
			user.Password = form.Password

			fmt.Printf("更新用户数据 %d", user.Id)
			_, err = engine.ID(user.Id).Update(user)
			if err != nil {
				print(err)
				c.JSON(200, gin.H{
					"code":    1,
					"message": "系统异常",
					"data":    "",
				})
				return
			}
			c.JSON(200, gin.H{
				"code":    0,
				"message": "密码重置成功",
				"data":    "",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "账号 邮箱不匹配",
				"data":    "",
			})
			return
		}
		//发送注册邮件
	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "create user error",
	})
}

type resetpasswordForm struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
