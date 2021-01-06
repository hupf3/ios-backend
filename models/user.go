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
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Gender string `json:"gender"`
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
	user.Username = ""
	user.UserID = -1
	user.Gender = ""
	user.Email = ""
	user.Phone = ""
	row := db.QueryRow("SELECT password FROM user where user_id = ?", userID)
	err := row.Scan(&user.Password)

	if err != nil {
		fmt.Printf("Query user failed, err:%v", err)
		return errors.New("用户不存在！")
	}

	if user.Password != password {
		return errors.New("密码不正确！")
	}

	// CurrentUserID = user.UserID
	return nil
}

// GetUserByID 获取用户信息
func GetUserByID(userID int) (User, error) {
	u := new(User)
	row := db.QueryRow("SELECT user_id, username, email, phone, gender FROM user where user_id = ?", userID)
	err := row.Scan(&u.UserID, &u.Username, &u.Email, &u.Phone, &u.Gender)

	if err != nil {
		fmt.Printf("Query user failed, err:%v\n", err)
		return *u, errors.New("User does not exists")
	}

	return *u, nil
}

// UpdateUser 修改用户信息
func UpdateUser(u User) (User, error) {
	oldUser, _ := GetUserByID(u.UserID)
	if u.Username == "" {
		u.Username = oldUser.Username
	}
	if u.Password == "" {
		u.Password = oldUser.Password
	}
	if u.Email == "" {
		u.Email = oldUser.Email
	}
	if u.Phone == "" {
		u.Phone = oldUser.Phone
	}
	if u.Gender == "" {
		u.Gender = oldUser.Gender
	}

	// stmt, err := db.Prepare("UPDATE user SET username = ?, location = ?, week_time = ?, term_time = ?, symbol = ? WHERE course_id = ?")
	stmt, err := db.Prepare("UPDATE user SET username = ?, password = ?, email = ?, phone = ?, gender = ? WHERE user_id = ?")
	_, err = stmt.Exec(u.Username, u.Password, u.Email, u.Phone, u.Gender, u.UserID)

	if err != nil {
		fmt.Printf("Update user failed, err:%v", err)
		return User{}, errors.New("User does not exists")
	}

	return u, nil
}
