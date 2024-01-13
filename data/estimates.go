/*
File: estimates.go

Status Codes:
1 - Not Contacted
2 - Contacted
3 - Converted
4 - Rejected
*/
package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type EstimateRequest struct {
	ID                int    `json:"id" db:"id"`
	StatusId          int    `json:"status" db:"status"`
	Name              string `json:"name" db:"name"`
	Email             string `json:"email" db:"email"`
	Phone             string `json:"phone" db:"phone"`
	Description       string `json:"description" db:"description"`
	ContactStatusCode int    `json:"status_code" db:"status_code"`
}

func (e EstimateRequest) InsertEstimateRequest(db *sql.DB) error {
	// I dont care if it already exists
	_, err := db.Exec(
		"INSERT INTO estimate_requests (name, email, phone, _description, estimate_status_code) VALUES (?,?,?,?,?)",
		e.Name, e.Email, e.Phone, e.Description, e.ContactStatusCode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (e EstimateRequest) UpdateEstimateRequest(db *sql.DB) {

}
