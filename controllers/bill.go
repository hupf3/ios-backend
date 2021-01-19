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

// GetAllBillsByUserID 获取某人账单
func GetAllBillsByUserID(context *gin.Context) {
	param := context.Param("userID")
	userID, _ := strconv.Atoi(param)

	data, err := models.GetAllBillsByUserID(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "succeed", "data": data, "count": len(data)})
}

// GetAllBills 获取全部账单
func GetAllBills(context *gin.Context) {
	data := make([]models.Bill, 0)
	data, err := models.GetAllBills()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "succeed", "data": data, "count": len(data)})
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
