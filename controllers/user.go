package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hupf3/ios-backend/models"
)

// SignUp 注册
func SignUp(c *gin.Context) {
	var u models.User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    "binding error",
		})
		return
	}

	if err := models.InsertUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed", "msg": "注册成功！"})
}

// Login 登陆
func Login(c *gin.Context) {
	// 读取json格式
	var obj map[string]interface{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Next()

	userIDT := obj["user_id"].(float64)
	var userID int = int(userIDT)

	passwordT := obj["password"].(string)
	var password string = passwordT

	if err := models.QueryUser(userID, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed", "msg": "登录成功！"})
}
