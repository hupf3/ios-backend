package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hupf3/ios-backend/controllers"
)

// SetRouter 设置路由
func SetRouter() *gin.Engine {
	router := gin.Default()

	/********** 用户信息 **********/
	// 注册与登录
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	/********** 课程信息 **********/

	/********** 作业信息 **********/

	/********** 账单信息 **********/
	router.PUT("/bill", controllers.AddBill)                        // 增加账单信息
	router.DELETE("/bills/:billID", controllers.DeleteBillByBillID) // 删除账单信息
	router.GET("/bills/:billID", controllers.GetBillByBillID)       // 获取账单信息
	router.POST("/bills/:billID", controllers.UpdateBillByBillID)   // 修改账单信息

	return router
}
