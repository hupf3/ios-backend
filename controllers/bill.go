package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"


	"github.com/hupf3/ios-backend/models"
	// "github.com/KianKw/ios-backend/models"
	"github.com/gin-gonic/gin"
)

// AddBill 增加账单
func AddBill(c *gin.Context) {
	var b models.Bill

	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    "binding error",
		})
		return
	}

	if err := models.InsertBill(b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed"})
}

// DeleteBillByBillID 删除账单
func DeleteBillByBillID(c *gin.Context) {
	param := c.Param("billID")
	billID, _ := strconv.Atoi(param)

	if err := models.DeleteBill(billID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed"})
}

// GetBillByBillID 获取账单
func GetBillByBillID(c *gin.Context) {
	param := c.Param("billID")
	billID, _ := strconv.Atoi(param)

	b := new(models.Bill)
	var err error

	if b, err = models.QueryBill(billID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "succeed",
		"data":   *b,
	})
}

// GetAllBillsByUserID 获取某人的全部账单
func GetAllBillsByUserID(c *gin.Context) {
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

	tmp := obj["user_id"].(float64)
	var userID int = int(tmp)

	bills := make([]models.Bill, 0)

	if bills, err = models.GetAllBills(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "succeed", "data": bills, "count": len(bills)})
}

// UpdateBillByBillID 修改账单
func UpdateBillByBillID(c *gin.Context) {
	param := c.Param("billID")
	billID, _ := strconv.Atoi(param)

	var b models.Bill
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    "binding error",
		})
		return
	}

	if err := models.UpdateBill(billID, b.Money, b.Classify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed"})
}
