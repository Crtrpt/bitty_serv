package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

func uploadAvatar(c *gin.Context) {
	//TODO 上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	// Upload the file to specific dst.
	name := node.Generate().Base64() + ".png"
	c.SaveUploadedFile(file, os.Getenv("asset_avatar_dir")+name)

	c.JSON(200, gin.H{
		"code": 0,
		"data": os.Getenv("asset_avatar_url") + name,
	})
}

func uploadImage(c *gin.Context) {
	//TODO 上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	// Upload the file to specific dst.
	name := node.Generate().Base64() + ".png"
	c.SaveUploadedFile(file, os.Getenv("asset_image_dir")+name)

	c.JSON(200, gin.H{
		"code": 0,
		"data": os.Getenv("asset_image_url") + name,
	})
}

func uploadFile(c *gin.Context) {
	//TODO 上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	// Upload the file to specific dst.
	name := node.Generate().Base64() + ".png"
	c.SaveUploadedFile(file, os.Getenv("asset_file_dir")+name)

	c.JSON(200, gin.H{
		"code": 0,
		"data": os.Getenv("asset_file_url") + name,
	})
}
