package models

// User 用户
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

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
