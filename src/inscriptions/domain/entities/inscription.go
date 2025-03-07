package entities

type Inscription struct {
	ID        int    `json:"id"`
	CourseID  int    `json:"course_id"`
	StudentID int    `json:"student_id"`
	Status    string `json:"status"` 
}


