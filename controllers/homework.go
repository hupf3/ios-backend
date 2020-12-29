package controllers

import (
	"net/http"
	"strconv"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/zwx2000/ios-backend/models"
)

//Get方法的查询
func GetHomeworks(c *gin.Context) {
	h := Homework{}
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

//利用Get方法通过id查询
func GetHomeworkByID(c *gin.Context) {
    var result gin.H
    //c.Params方法可以获取到/homework/:hw_id中的id值
    hw_id := c.Param("Id")
    Id, err := strconv.Atoi(hw_id)
    if err != nil {
        log.Fatal(err)
    }
    //定义homework结构
    h := Homework{
        HomeworkID: Id,
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

//利用post方法新增数据
func AddNewHomework(c *gin.Context) {
    var h Homework
    err := c.Bind(&h)
    if err != nil {
        log.Fatal(err)
    }
    Id, err := h.AddHomework()
    fmt.Print("id=", Id)
    content := h.Content;
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Successfully add homework: %s", content),
    })
}

//利用DELETE请求方法通过id删除
func DeleteHomeworkByID(c *gin.Context) {
    id := c.Param("Id")

    Id, err := strconv.ParseInt(id, 10, 10)
    if err != nil {
        log.Fatalln(err)
    }
    h := Homework{HomeworkID: int(Id)}
        
    rows, err := h.DeleteHomework()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("delete rows ", rows)

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Successfully deleted homework: %s", id),
    })
}

//利用PUT请求方法修改作业内容信息
func UpdateHomeworkByID(c *gin.Context) {
  	cid := c.Param("id")
  	id, err := strconv.Atoi(cid)

  	h := Homework{HomeworkID: id}
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