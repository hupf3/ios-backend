package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zwx2000/ios-backend/controllers"
)

// SetRouter 设置路由
func SetRouter() *gin.Engine {
	router := gin.Default()

	/********** 用户信息 **********/
	// 注册与登录
	// router.POST("/signup", controller.SignUp)
	// router.POST("/login", controller.Login)

	/********** 课程信息 **********/

	/********** 作业信息 **********/
	router.PUT("/homework/:id", controllers.UpdateHomeworkByID)
	router.DELETE("/homework/:id", controllers.DeleteHomeworkByID)
	router.GET("/homework", controllers.GetHomeworks)
	router.GET("/homework/:id", controllers.GetHomeworkByID)
	router.POST("/homework", controllers.AddNewHomework)

	/********** 账单信息 **********/
	router.PUT("/bill", controllers.AddBill)                        // 增加账单信息
	router.DELETE("/bills/:billID", controllers.DeleteBillByBillID) // 删除账单信息
	router.GET("/bills/:billID", controllers.GetBillByBillID)       // 获取账单信息
	router.POST("/bills/:billID", controllers.UpdateBillByBillID)   // 修改账单信息

	return router
}
