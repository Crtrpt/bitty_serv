package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func signupEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

var engine *xorm.Engine
f, err := os.Create("sql.log")

func Router() http.Handler {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	e := gin.New()
	e.Use(gin.Recovery())
	v1 := e.Group("/v1/auth")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/signup", signupEndpoint)
	}
	return e
}
