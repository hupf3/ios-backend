package models

import (
	"database/sql"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// 数据库的信息
const (
	USERNAME = "root"      // 用户名
	PASSWORD = "huq35999"  // 密码
	DATABASE = "mosad"     // 数据库名字
	PORT     = "3306"      // 端口
	IP       = "127.0.0.1" // IP地址
)

var db *sql.DB

// Connect 连接数据库
func Connect() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, IP, PORT, DATABASE)
	var err error
	db, err = sql.Open("mysql", connectInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect success!")
}
