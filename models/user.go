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
		return errors.New("User exists")
	}

	return nil
}
