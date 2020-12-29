package controllers

import (
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
			"msg":  "binding error",
		})
		return
	}

	if err := models.InsertUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed"})
}

// Login 登陆
func Login(c *gin.Context) {

}
