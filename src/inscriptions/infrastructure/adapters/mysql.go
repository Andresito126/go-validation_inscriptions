package adapters

import (
    "database/sql"
    "fmt"
)

type MySQLAdapter struct {
    db *sql.DB
}

func NewMySQLAdapter(db *sql.DB) *MySQLAdapter {
    return &MySQLAdapter{db: db}
}

func (a *MySQLAdapter) Validate(inscriptionID int) (string, error) {
    var courseID, studentID int

   // get courseid y studentid
    err := a.db.QueryRow(`
        SELECT course_id, student_id 
        FROM inscriptions 
        WHERE id = ?
    `, inscriptionID).Scan(&courseID, &studentID)
    if err != nil {
        return "", err
    }

    // primer validacion de saber esta en el mismo curso
    var existingID int
    err = a.db.QueryRow(`
        SELECT id FROM inscriptions 
        WHERE course_id = ? AND student_id = ? AND status = 'aceptada'
    `, courseID, studentID).Scan(&existingID)
    if err == nil {
        fmt.Println(" Estudiante ya inscrito en este curso, rechazando inscripción ID:", inscriptionID)
        return "rechazada", nil 
    } else if err != sql.ErrNoRows {
        fmt.Println("Error en consulta SQL:", err)
        return "", err
    }
    

    // verificar cupo
    var availableSlots int
    err = a.db.QueryRow(`
        SELECT available_slots 
        FROM courses 
        WHERE id = ?
    `, courseID).Scan(&availableSlots)
    if err != nil {
        return "", err
    }

    if availableSlots <= 0 {
        return "rechazada", nil
    }

     // quita uno al cruso
    //verificacion de cupo
    _, err = a.db.Exec(`
        UPDATE courses 
        SET available_slots = available_slots - 1 
        WHERE id = ?
    `, courseID)
    if err != nil {
        return "", err
    }

    return "aceptada", nil
}


func (a *MySQLAdapter) UpdateStatus(inscriptionID int, status string) error {
    _, err := a.db.Exec(`
        UPDATE inscriptions 
        SET status = ? 
        WHERE id = ?
    `, status, inscriptionID)
    return err
}
