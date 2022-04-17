package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bitty/middleware"
	"bitty/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine
var err error

func Router() http.Handler {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine, err = xorm.NewEngine("mysql", os.Getenv("db"))
	engine.ShowSQL(true)

	engine.SetTableMapper(names.SnakeMapper{})
	engine.SetColumnMapper(names.SnakeMapper{})

	engine.Sync2(new(model.User))
	engine.Sync2(new(model.Endpoint))

	var rows, _ = engine.Query("select version() `version`")
	fmt.Printf("\n=========================================================\n\n")

	fmt.Printf("DBVersion: %s", rows[0]["version"])

	fmt.Printf("\n=========================================================\n\n")
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
	return e
}

func login(c *gin.Context) {
	var form PostLogin
	if c.BindJSON(&form) == nil {
		var u, err = engine.Get(&model.User{Account: form.Account, Password: form.Password})
		if err != nil {
			fmt.Printf("ERROR:%s", err)
		}
		if u {
			c.JSON(200, gin.H{
				"code": 0,
				"data": "xxxxxxxxxxxxxxxxx",
			})
			return
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
		_, err = engine.Insert(user)
		if err == nil {
			c.JSON(200, gin.H{
				"code": 0,
				"data": "xxxxxxxxxxxxxxxxx",
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
