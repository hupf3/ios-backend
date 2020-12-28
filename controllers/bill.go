package controllers

import (
	"github.com/gin-gonic/gin"
)

// AddBill 增加账单
func AddBill(c *gin.Context) {

}

// DeleteBillByBillID 删除账单
func DeleteBillByBillID(c *gin.Context) {

}

// GetBillByBillID 获取账单
func GetBillByBillID(c *gin.Context) {
	billID := c.Param("billID")
	println(billID)
}

// UpdateBillByBillID 修改账单
func UpdateBillByBillID(c *gin.Context) {

}
