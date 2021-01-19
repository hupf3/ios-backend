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
	IsFinished int    `json:"is_finished"`
}

// AddHomework 添加作业
func (h Homework) AddHomework() (ID int, err error) {
	stmt, err := db.Prepare("INSERT INTO homework(hw_id, user_id, course_id, content, deadline, is_finished) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}

	// 执行插入操作
	rs, err := stmt.Exec(h.HomeworkID, h.UserID, h.CourseID, h.Content, h.Deadline, h.IsFinished)
	if err != nil {
		return
	}

	// 返回插入的id
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	// 将id类型转换
	ID = int(id)

	defer stmt.Close()
	return
}

// DeleteHomework 根据id删除作业
func (h Homework) DeleteHomework() (rows int, err error) {
	stmt, err := db.Prepare("DELETE FROM homework WHERE hw_id=?")
	if err != nil {
		log.Fatalln(err)
	}

	rs, err := stmt.Exec(h.HomeworkID)
	if err != nil {
		log.Fatalln(err)
	}
	// 删除的行数
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
	return
}

// GetHomework 根据id获取对应的作业信息
func (h Homework) GetHomework() (homework Homework, err error) {
	row := db.QueryRow("SELECT * FROM homework WHERE hw_id=?", h.HomeworkID)
	err = row.Scan(&homework.HomeworkID, &homework.UserID, &homework.CourseID, &homework.Content, &homework.Deadline, &homework.IsFinished)
	if err != nil {
		fmt.Println("fail to get!")
		return
	}
	return
}

// GetAllHomework 获取所有作业信息
func (h Homework) GetAllHomework() (homeworks []Homework, err error) {
	rows, err := db.Query("SELECT * FROM homework")
	if err != nil {
		return
	}
	for rows.Next() {
		var homework Homework
		// 遍历表中所有行的信息
		rows.Scan(&homework.HomeworkID, &homework.UserID, &homework.CourseID, &homework.Content, &homework.Deadline, &homework.IsFinished)
		// 将person添加到persons中
		homeworks = append(homeworks, homework)
	}
	//最后关闭连接
	defer rows.Close()
	return
}

// GetUnfinishedHomeworkByUser 获取某人没有完成的作业
func GetUnfinishedHomeworkByUser(userID int) ([]Homework, error) {
	homeworks := make([]Homework, 0)
	rows, err := db.Query("SELECT * FROM homework WHERE user_id = ? AND is_finished = 0 ORDER BY deadline", userID)
	if err != nil {
		fmt.Printf("Query homeworks failed, err:%v", err)
		return nil, err
	}
	for rows.Next() {
		var homework Homework
		if err = rows.Scan(&homework.HomeworkID, &homework.UserID, &homework.CourseID, &homework.Content, &homework.Deadline, &homework.IsFinished); err != nil {
			fmt.Printf("Scan homework failed, err:%v", err)
			return nil, err
		}
		homeworks = append(homeworks, homework)
	}
	return homeworks, nil
}

// GetHomeworksByUserAndCourse 获取某人某课程作业
func GetHomeworksByUserAndCourse(userID int, courseID int) ([]Homework, error) {
	homeworks := make([]Homework, 0)
	rows, err := db.Query("SELECT * FROM homework where user_id = ? AND course_id = ? ORDER BY is_finished, deadline", userID, courseID)
	if err != nil {
		fmt.Printf("Query homeworks failed, err:%v", err)
		return nil, err
	}
	for rows.Next() {
		var homework Homework
		if err = rows.Scan(&homework.HomeworkID, &homework.UserID, &homework.CourseID, &homework.Content, &homework.Deadline, &homework.IsFinished); err != nil {
			fmt.Printf("Scan homework failed, err:%v", err)
			return nil, err
		}
		homeworks = append(homeworks, homework)
	}
	return homeworks, nil
}

// UpdateHomework 更新作业内容
func (h Homework) UpdateHomework() (err error) {
	stmt, err := db.Prepare("UPDATE homework SET is_finished=? WHERE hw_id=?")
	if err != nil {
		log.Fatalln(err)
	}
	h.IsFinished = 1
	rs, err := stmt.Exec(h.IsFinished, h.HomeworkID)
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
