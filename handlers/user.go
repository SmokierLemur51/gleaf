package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SmokierLemur51/gleaf/data"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID             int    `json:"id" db:"id"`
	Email          string `json:"email" db:"id"`
	Username       string `json:"username" db:"username"`
	PassHash       string `json:"passhash" db:"passhash"`
	ClearanceLevel string
	ClearanceID    int    `json:"clearance" db:"clearance_level"`
	SessionID      string `json:"session_id" db:"session_id"`
}

// not quite done ...
func LoadUser(db *sql.DB, email string) (User, error) {
	var u User
	err := db.QueryRow("SELECT id, email, username, passhash, clearance_level, session_id FROM users WHERE email = ?",
		email,
	).Scan(&u.ID, &u.Email, &u.Username, &u.PassHash, &u.ClearanceID, &u.SessionID)
	if err == sql.ErrNoRows {
		// doesnt exist
		return u, nil
	} else if err != nil {
		// err in query
		return u, err
	}
	return u, nil
}

func (u User) InsertUser(db *sql.DB) {
	var execute bool
	var err error
	execute, err = data.CheckExistence(db, "users", "email", u.Email)
	if err != nil {
		log.Println(err)
		return
	}

	// remember, the check existing returns true if the category already exists, so it skips
	switch execute {
	case false:
		if u.ClearanceID == 0 {
			u.ClearanceID = data.FindDatabaseID(db, "clearance", "clearance_level", u.ClearanceLevel)
		}
		_, err := db.Exec(
			"INSERT INTO users (email, username, password_hash, clearance_level) VALUES (?,?,?,?)",
			u.Email, u.Username, u.PassHash, u.ClearanceID,
		)
		if err != nil {
			log.Fatal(err)
		}
	case true:
		fmt.Printf("Email <%s> exists with username <%s> already exists.\n", u.Email, u.Username)
	}
}

// not done
func (u User) UpdateUserInformation(db *sql.DB) {}
