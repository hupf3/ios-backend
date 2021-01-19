package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Upload 上传头像
func Upload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := fmt.Sprintf("./imgs/%s", file.Filename)
	// 上传文件到指定的 dst 。
	c.SaveUploadedFile(file, dst)

	c.JSON(http.StatusOK, gin.H{"status": "succeed", "data": dst[1:]})
}

func ShowImage(context *gin.Context) {
	imgurl := context.Param("imgurl")
	context.File(imgurl)
}
