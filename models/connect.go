package models

import (
	"database/sql"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// 数据库的信息
const (
	USERNAME = "root"
	PASSWORD = "huq35999"
	DATABASE = "iosPro"
)

var db *sql.DB

// Connect 连接数据库
func Connect() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, DATABASE)
	var err error
	db, err = sql.Open("mysql", connectInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect success!")
}
