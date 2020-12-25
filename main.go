package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库的信息
const (
	USERNAME = "root"
	PASSWORD = "huq35999"
	DATABASE = "iosPro"
)

var db *sql.DB

// CreateUserTable 创建用户表
func CreateUserTable() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		user_id INT NOT NULL AUTO_INCREMENT,
		user_name VARCHAR(32) UNIQUE,
		password VARCHAR(32),
		bio VARCHAR(128) DEFAULT '',
		PRIMARY KEY (user_id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Create userTable failed!", err)
		return
	}
}

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

func main() {
	Connect()
	CreateUserTable()

	// 断开数据库连接
	defer db.Close()
}
