package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func uploadImage(c *gin.Context) {
	//TODO 上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	// Upload the file to specific dst.
	name := node.Generate().Base64() + ".png"
	c.SaveUploadedFile(file, os.Getenv("asset_avatar_dir")+name)

	c.JSON(200, gin.H{
		"code": 0,
		"data": os.Getenv("asset_url") + name,
	})
}
