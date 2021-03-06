package main

import (
	"github.com/hupf3/ios-backend/models"
	"github.com/hupf3/ios-backend/routes"
)

func main() {
	// 连接数据库
	models.Connect()

	// 设置路由
	router := routes.SetRouter()
	router.Run(":9090")

	// 断开数据库连接
	// defer db.Close()
}
