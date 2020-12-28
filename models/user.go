package models

// User 用户
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	// Avatar   string `json:"avatar"`
}
