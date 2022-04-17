package api

import (
	"github.com/gin-gonic/gin"
	"bitty/model"
)

func list(c *gin.Context) {
	var userId =c.Request.URL.Query().Get("userId")

	var list []model.Endpoint

	err := engine.Where("userId = ?", userId).Find(&list)
	if(err!=nil){
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
	})
}

func search(c *gin.Context) {
	var keywords =c.Request.URL.Query().Get("keywords")

	var list []model.User

	var query=engine.Cols("nick_name","user_id").Where("nick_name like ?",keywords+"%").Limit(10)
	query.Find(&list)

	
	c.JSON(200, gin.H{
		"code": 0,
		"data": list,
	})
	return
}

func add(c *gin.Context) {

	var userId =c.Request.URL.Query().Get("userId")

	var user = &model.User{UserId:userId}

	has ,err := engine.Get(user)
	if(!has || err!=nil){
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "user not found",
		})
		return
	}

	endpoint := new(model.Endpoint)

	endpoint.Name=user.NickName;
	endpoint.UserId=user.UserId;
	endpoint.MemberAmount=0;
	endpoint.Scenes=0;
	endpoint.EndpointId=node.Generate().Base64();
	_, err = engine.Insert(endpoint)
	
	c.JSON(200, gin.H{
		"code": 0,
		"data": endpoint.EndpointId,
	})
	return
}
