/*
Some useful notes for JWT

Ok to put user credentials in token? - Not sensitive info
https://stackoverflow.com/questions/42652695/is-it-ok-to-store-user-credentials-in-the-jwt
*/

package handlers

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/SmokierLemur51/gleaf/data"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
)

type CustomClaims struct {
	Username string
	Password string
	jwt.RegisteredClaims
}

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

// not done
func (u User) VerifyCredentials(db *sql.DB, f LoginForm) bool { return false }

func VerifyCredentials(db *sql.DB, email, pass string) (User, error) {
	var u User
	var exists bool
	exists, err := data.CheckExistence(db, "users", "email", email)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		// compare hash

		// load user
	} else {
		// it doesnt exist
		return u, nil
	}

	return u, nil
}

// not done
func CreateToken(sessionID, username string) {}

// not done
func GenerateSessionID(length int) string { return "" }

func GenerateKey(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}

		randomString[i] = charset[randomIndex.Int64()]
	}
	fmt.Printf("\n\t%s\n\n", string(randomString))
	return string(randomString), nil
}

func SetKeyEnvironment(key string, length int) error {
	val, err := GenerateKey(length)
	if err != nil {
		return err
	}

	if err := os.Setenv(key, val); err != nil {
		return err
	}

	return nil
}

func LoadEnvKey(key string) string {
	return string(os.Getenv(key))
}
