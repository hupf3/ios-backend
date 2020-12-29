package models

import (
	"errors"
	"fmt"
)

// Course 课程
type Course struct {
	CourseID   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	Location   string `json:"location"`
	WeekTime   string `json:"week_time"`
	TermTime   string `json:"term_time"`
	Symbol     int    `json:"symbol"`
}

// CreateCourse 创建一个课程
func CreateCourse(c Course) (Course, error) {
	stmt, err := db.Prepare("INSERT INTO course(course_id, course_name, location, week_time, term_time, symbol) values(?,?,?,?,?,?)")
	_, err = stmt.Exec(c.CourseID, c.CourseName, c.Location, c.WeekTime, c.TermTime, c.Symbol)
	if err != nil {
		fmt.Printf("Insert course failed, err:%v", err)
		return Course{}, errors.New("Course exists")
	}
	return c, nil
}

// DeleteCourseByID 通过 ID 删除一个课程
func DeleteCourseByID(CourseID int) error {
	_, err := db.Exec("DELETE FROM course WHERE course_id = ?", CourseID)
	if err != nil {
		fmt.Printf("Delete course failed, err:%v", err)
		return errors.New("Course does not exists")
	}
	return nil
}

// GetCourseByID 通过 ID 获取一个课程
func GetCourseByID(courseID int) (Course, error) {
	c := new(Course)
	row := db.QueryRow("SELECT * FROM course where course_id = ?", courseID)
	err := row.Scan(&c.CourseID, &c.CourseName, &c.Location, &c.WeekTime, &c.TermTime, &c.Symbol)

	if err != nil {
		fmt.Printf("Query course failed, err:%v", err)
		return *c, errors.New("Course does not exists")
	}

	return *c, nil
}

// UpdateCourse 通过 ID 更新一个课程
func UpdateCourse(c Course) (Course, error) {
	oldCourse, _ := GetCourseByID(c.CourseID)
	if c.CourseName == "" {
		c.CourseName = oldCourse.CourseName
	}
	if c.Location == "" {
		c.Location = oldCourse.Location
	}
	if c.WeekTime == "" {
		c.WeekTime = oldCourse.WeekTime
	}
	if c.TermTime == "" {
		c.TermTime = oldCourse.TermTime
	}
	if c.CourseName == "" {
		c.Symbol = oldCourse.Symbol
	}

	stmt, err := db.Prepare("UPDATE course SET course_name = ?, location = ?, week_time = ?, term_time = ?, symbol = ? WHERE course_id = ?")
	_, err = stmt.Exec(c.CourseName, c.Location, c.WeekTime, c.TermTime, c.Symbol, c.CourseID)

	if err != nil {
		fmt.Printf("Update course failed, err:%v", err)
		return Course{}, errors.New("Course does not exists")
	}

	return c, nil
}
