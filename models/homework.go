package models

// Homework 作业
type Homework struct {
	HomeworkID int    `json:"hw_id"`
	UserID     int    `json:"user_id"`
	CourseID   int    `json:"course_id"`
	Content    string `json:"content"`
	Deadline   string `json:"deadline"`
}
