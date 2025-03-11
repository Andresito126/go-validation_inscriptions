package models

type Inscription struct {
    ID        int    `json:"id"`
    StudentID int    `json:"student_id"`
    CourseID  int    `json:"course_id"`
    Status    string `json:"status"`
}
