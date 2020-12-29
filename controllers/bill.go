package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hupf3/ios-backend/models"
)

// AddBill 增加账单
func AddBill(c *gin.Context) {
	var b models.Bill

	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "binding error",
		})
		return
	}

	if err := models.InsertBill(b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
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
			"error":  err.Error(),
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
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   *b,
	})
}

// UpdateBillByBillID 修改账单
func UpdateBillByBillID(c *gin.Context) {
	param := c.Param("billID")
	billID, _ := strconv.Atoi(param)

	var b models.Bill
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "binding error",
		})
		return
	}

	if err := models.UpdateBill(billID, b.Money, b.Tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succeed"})
}
