package main

import (
	"github.com/hupf3/ios-backend/models"
	// "github.com/hupf3/ios-backend/route"
	// "github.com/hupf3/ios-backend/service"
)

// // CreateUserTable 创建用户表
// func CreateUserTable() {
// 	sql := `CREATE TABLE IF NOT EXISTS users(
// 		user_id INT NOT NULL AUTO_INCREMENT,
// 		user_name VARCHAR(32) UNIQUE,
// 		password VARCHAR(32),
// 		bio VARCHAR(128) DEFAULT '',
// 		PRIMARY KEY (user_id)
// 		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

// 	// 执行建表
// 	if _, err := db.Exec(sql); err != nil {
// 		fmt.Println("Create userTable failed!", err)
// 		return
// 	}
// }

func main() {
	models.Connect()

	// 断开数据库连接
	// defer db.Close()
}
