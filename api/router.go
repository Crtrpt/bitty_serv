package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() http.Handler {

	Init()
	e := gin.New()
	e.StaticFS("./upload", http.Dir("./upload"))
	e.MaxMultipartMemory = 8 << 20
	e.Use(gin.Recovery())
	e.Use(CORSMiddleware())
	v1 := e.Group("/api/v1/auth")
	{
		v1.POST("/login", login)
		v1.POST("/signup", signup)
		v1.POST("/sendcode", sendCode)
		v1.POST("/resetpassword", resetpassword)
	}
	v2 := e.Group("/api/v1/contact")
	{
		v2.Use(TokenMiddleware())
		v2.POST("/add", addContact)
		v2.POST("/remove", removeContact)
		v2.GET("/list", list)
		v2.GET("/info", infoContact)
		v2.GET("/search", search)
	}
	v3 := e.Group("/api/v1/user")
	{
		v3.Use(TokenMiddleware())
		v3.GET("/profile", profile)
		v3.GET("/session", UserSession)
		v3.POST("/save", save)
	}
	msg := e.Group("/api/v1/msg")
	{
		msg.Use(TokenMiddleware())
		msg.GET("/unreadMessage", unreadMessage)
		msg.GET("/allMessage", allMessage)
		msg.POST("/action", messageAction)
	}
	asset := e.Group("/api/v1/asset")
	{
		asset.Use(TokenMiddleware())
		asset.POST("/uploadAvatar", uploadAvatar)
		asset.POST("/uploadImage", uploadImage)
		asset.POST("/uploadFile", uploadFile)
	}

	session := e.Group("/api/v1/session")
	{
		session.Use(TokenMiddleware())
		session.POST("/create", SessionCreate)
		session.POST("/toggle_suspend", SessionSuspend)
		session.GET("/list", SessionList)
		session.GET("/info", SessionInfo)
		session.GET("/profile", SessionProfile)
	}

	chat := e.Group("/api/v1/chat")
	{
		chat.Use(TokenMiddleware())
		chat.POST("/sendMsg", sendMsg)
	}

	group := e.Group("/api/v1/group")
	{
		group.Use(TokenMiddleware())
		group.POST("/create", GroupCreate)
		group.GET("/list", GroupList)
	}
	return e
}
