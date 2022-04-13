package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary  login get auth code
// @Schemes
// @Description  login
// @Tags         example
// @Accept       json
// @Produce      json
// @Success      200  {string}  8234625472354763285476324
// @Router       /auth/login [get]
func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func infoEndpoint(c *gin.Context) {
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
var err error

func Router() http.Handler {

	engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

	e := gin.New()
	e.Use(gin.Recovery())
	v1 := e.Group("/v1/auth")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/signup", signupEndpoint)
		v1.GET("/info", infoEndpoint)
	}
	return e
}
