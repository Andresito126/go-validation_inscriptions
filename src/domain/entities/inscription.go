package entities


type Inscription struct {
	ID        int              `json:"id"`
	StudentID int              `json:"student_id"`
	CourseID  int              `json:"course_id"`
	Status    InscriptionStatus  `json:"status"`
}



type InscriptionStatus string

const (
	Pending   InscriptionStatus = "pending"
	Confirmed InscriptionStatus = "confirmed"
    Rejected  InscriptionStatus = "rejected"
)




