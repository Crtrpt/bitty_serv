package auth

import (
	"net/http"

	"bitty/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine
var err error

func Router() http.Handler {

	engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(middleware.CORSMiddleware())
	v1 := e.Group("/api/v1/auth")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/signup", signupEndpoint)
	}
	return e
}

// @BasePath /api/v1

// @Summary  用户登录
// @Schemes
// @Description  用户登录
// @Tags         授权
// @Accept       json
// @Produce      json
// @Success      200  {string}  8234625472354763285476324
// @Router       /auth/login [post]
// @Param        account   formData  string  true  "account"
// @Param        password  formData  string  true  "password"
func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

// @Summary  用户注册
// @Schemes
// @Description  用户登录
// @Tags         授权
// @Accept       json
// @Produce      json
// @Success      200  {string}  8234625472354763285476324
// @Router       /auth/signup [post]
// @Param        account   query  string  true  "account"
// @Param        password  query  string  true  "password"
func signupEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}
