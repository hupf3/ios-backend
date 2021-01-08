package routes

import (
	// "github.com/hupf3/ios-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/hupf3/ios-backend/controllers"
)

// SetRouter 设置路由
func SetRouter() *gin.Engine {
	router := gin.Default()

	/********** 用户信息 **********/
	// 注册与登录
	router.POST("/signup", controllers.SignUp)            // 注册
	router.POST("/login", controllers.Login)              // 登录
	router.GET("/users/:userID", controllers.GetUserByID) // 获取用户信息
	router.PUT("/users/:userID", controllers.UpdateUser)  // 修改用户信息

	/********** 课程信息 **********/
	router.POST("/courses", controllers.CreateCourse)                 // 增加课程信息
	router.DELETE("/courses/:courseID", controllers.DeleteCourseByID) // 删除课程信息
	router.GET("/courses/:courseID", controllers.GetCourseByID)       // 获取课程信息
	router.GET("/courses", controllers.GetAllCourses)                 // 获取所有课程
	router.PUT("/courses/:courseID", controllers.UpdateCourseByID)    // 修改课程信息

	/********** 作业信息 **********/
	router.POST("/homework", controllers.AddNewHomework)             // 增加作业信息
	router.DELETE("/homework/:hwID", controllers.DeleteHomeworkByID) // 删除作业信息
	router.GET("/homework", controllers.GetHomeworks)                // 获取全部作业
	router.GET("/homework/:hwID", controllers.GetHomeworkByID)       // 获取作业信息
	router.PUT("/homework/:hwID", controllers.UpdateHomeworkByID)    // 判断作业是否完成
	// router.GET("/users/:userID/courses/:courseID/homeworks", controllers.GetHomeworksByUserAndCourse) // 获取某人某课程作业
	// router.GET("/users/:userID/homeworks/unfinished", controllers.GetUnfinishedHomeworkByUser)        // 获取未完成作业信息

	/********** 账单信息 **********/
	router.POST("/bills", controllers.AddBill)                          // 增加账单信息
	router.DELETE("/bills/:billID", controllers.DeleteBillByBillID)     // 删除账单信息
	router.GET("/bills/:billID", controllers.GetBillByBillID)           // 获取账单信息
	router.GET("/bills", controllers.GetAllBills)                       // 获取全部账单
	router.PUT("/bills/:billID", controllers.UpdateBillByBillID)        // 修改账单信息
	router.GET("/users/:userID/bills", controllers.GetAllBillsByUserID) // 修改某人的账单信息

	return router
}
