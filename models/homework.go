package models

import (
	"fmt"
	"log"
)
// Homework 作业
type Homework struct {
	HomeworkID int    `json:"hw_id"`
	UserID     int    `json:"user_id"`
	CourseID   int    `json:"course_id"`
	Content    string `json:"content"`
	Deadline   string `json:"deadline"`
}

//添加作业
func (h Homework) AddHomework() (Id int, err error) {
    stmt, err := db.Prepare("INSERT INTO homework(hw_id, user_id, course_id, content) VALUES (?, ?, ?, ?)")
    if err != nil {
        return
    }

    //执行插入操作
    rs, err := stmt.Exec(h.HomeworkID, h.UserID, h.CourseID, h.Content)
    if err != nil {
        return
    }

    //返回插入的id
    id, err := rs.LastInsertId()
    if err != nil {
        log.Fatalln(err)
    }

    //将id类型转换
    Id = int(id)

    defer stmt.Close()
    return
}

//根据id删除作业
func (h Homework) DeleteHomework() (rows int, err error) {
    stmt, err := db.Prepare("DELETE FROM homework WHERE hw_id=?")
    if err != nil {
        log.Fatalln(err)
    }

    rs, err := stmt.Exec(h.HomeworkID)
    if err != nil {
        log.Fatalln(err)
    }
    //删除的行数
    row, err := rs.RowsAffected()
    if err != nil {
        log.Fatalln(err)
    }
    defer stmt.Close()
    rows = int(row)
    return
}

//根据id获取对应的作业信息
func (h Homework) GetHomework() (homework Homework, err error) {
    row := db.QueryRow("SELECT * FROM homework WHERE hw_id=?", h.HomeworkID)
    err = row.Scan(&homework.HomeworkID, &homework.UserID, &homework.CourseID, &homework.Content, &homework.Deadline)
    if err != nil {
    	fmt.Println("fail to get!")
        return
    }
    return
}

//获取所有作业信息
func (h Homework) GetAllHomework() (homeworks []Homework, err error) {
    rows, err := db.Query("SELECT * FROM homework")
    if err != nil {
        return
    }
    for rows.Next() {
        var homework Homework
        //遍历表中所有行的信息
        rows.Scan(&homework.HomeworkID, &homework.UserID, &homework.CourseID, &homework.Content, &homework.Deadline)
        //将person添加到persons中
        homeworks = append(homeworks, homework)
    }
    //最后关闭连接
    defer rows.Close()
    return
}

//更新作业内容
func (h Homework) UpdateHomework() (err error) {
  	stmt, err := db.Prepare("UPDATE homework SET content=? WHERE hw_id=?")
	if err != nil {
   		log.Fatalln(err)
  	}

  	rs, err := stmt.Exec(h.Content, h.HomeworkID)
  	if err != nil {
   		log.Fatalln(err)
  	}

  	_, err = rs.RowsAffected()
  	if err != nil {
   		log.Fatalln(err)
  	}

	defer stmt.Close()
  	return
}