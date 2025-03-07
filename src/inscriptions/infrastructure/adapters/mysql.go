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
    // verifica si hay cupo 
    var availableSlots int
    err := a.db.QueryRow(`
        SELECT available_slots 
        FROM courses 
        WHERE id = (SELECT course_id FROM inscriptions WHERE id = ?)
    `, inscriptionID).Scan(&availableSlots)
    if err != nil {
        return "", err
    }

    if availableSlots <= 0 {
        return "rechazada", nil
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