package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hupf3/ios-backend/models"
)

// GetHomeworks Get方法的查询
func GetHomeworks(c *gin.Context) {
	h := models.Homework{}
	homeworks, err := h.GetAllHomework()
	if err != nil {
		log.Fatal(err)
	}
	//H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"result": homeworks,
		"count":  len(homeworks),
	})
}

// GetHomeworkByID 利用Get方法通过id查询
func GetHomeworkByID(c *gin.Context) {
	var result gin.H
	// c.Params方法可以获取到/homework/:hw_id中的id值
	hwID := c.Param("hwID")
	ID, err := strconv.Atoi(hwID)
	if err != nil {
		log.Fatal(err)
	}
	// 定义homework结构
	h := models.Homework{
		HomeworkID: ID,
	}
	homework, err := h.GetHomework()
	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": homework,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

// AddNewHomework 利用post方法新增数据
func AddNewHomework(c *gin.Context) {
	var h models.Homework
	err := c.Bind(&h)
	if err != nil {
		log.Fatal(err)
	}
	ID, err := h.AddHomework()
	fmt.Print("id=", ID)
	content := h.Content
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully add homework: %s", content),
	})
}

// DeleteHomeworkByID 利用DELETE请求方法通过id删除
func DeleteHomeworkByID(c *gin.Context) {
	id := c.Param("hwID")

	ID, err := strconv.ParseInt(id, 10, 10)
	if err != nil {
		log.Fatalln(err)
	}
	h := models.Homework{HomeworkID: int(ID)}

	rows, err := h.DeleteHomework()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("delete rows ", rows)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted homework: %s", id),
	})
}

// UpdateHomeworkByID 利用PUT请求方法修改作业内容信息
func UpdateHomeworkByID(c *gin.Context) {
	cid := c.Param("hwID")
	id, err := strconv.Atoi(cid)

	h := models.Homework{HomeworkID: id}
	err = c.Bind(&h)
	if err != nil {
		log.Fatalln(err)
	}

	err = h.UpdateHomework()
	msg := fmt.Sprintf("Update homework %d successfully", h.HomeworkID)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
