package main

import (
	"github.com/hupf3/ios-backend/models"
	// "github.com/hupf3/ios-backend/route"
	// "github.com/hupf3/ios-backend/service"
)

func main() {
	models.Connect()

	// 断开数据库连接
	// defer db.Close()
}
