// src/api2/infrastructure/adapters/mysql.go
package adapters

import (
    "database/sql" 
)

type MySQLAdapter struct {
    db *sql.DB
}

func NewMySQLAdapter(db *sql.DB) *MySQLAdapter {
    return &MySQLAdapter{db: db}
}

func (a *MySQLAdapter) Validate(inscriptionID int) (string, error) {
    // get courseid
    var courseID int
    err := a.db.QueryRow(`
        SELECT course_id 
        FROM inscriptions 
        WHERE id = ?
    `, inscriptionID).Scan(&courseID)
    if err != nil {
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