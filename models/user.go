package models

import (
	"errors"
	"fmt"
)

// User 用户
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	// Avatar   string `json:"avatar"`
}

// InsertUser 添加用户
func InsertUser(u User) error {
	_, err := db.Exec("INSERT INTO user(user_id, username, password) values(?, ?, ?)", u.UserID, u.Username, u.Password)
	if err != nil {
		fmt.Printf("Insert user failed, err:%v", err)
		return errors.New("该用户已经注册！")
	}

	return nil
}

// QueryUser 查询用户信息
func QueryUser(userID int, password string) error {
	user := new(User)
	row := db.QueryRow("SELECT * FROM user where user_id = ?", userID)
	err := row.Scan(&user.UserID, &user.Username, &user.Password)

	if err != nil {
		fmt.Printf("Query user failed, err:%v", err)
		return errors.New("用户不存在！")
	}

	if user.Password != password {
		return errors.New("密码不正确！")
	}

	CurrentUserID = user.UserID
	return nil
}
