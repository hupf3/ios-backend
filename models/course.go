package models

// Course 课程
type Course struct {
	CourseID   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	Location   string `json:"location"`
	WeekTime   string `json:"week_time"`
	TermTime   string `json:"term_time"`
	Symbol     int    `json:"symbol"`
}
