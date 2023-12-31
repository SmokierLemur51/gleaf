package data

import (
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// just need to fix err handling
func CheckExistence(db *sql.DB, table, column, item string) (bool, error) {
	// returns true if it exists
	var count int
	stmt, err := db.Prepare(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(item)
	if err != nil {
		return true, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false, err
		}
	}
	if count > 0 {
		return true, err
	}
	// if not

	return false, nil
}

// not done needs to also return error
func FindDatabaseID(db *sql.DB, table, column, item string) (int, error) {
	stmt, err := db.Prepare(fmt.Sprintf("SELECT id FROM %s WHERE %s = ?", table, column))
	if err != nil {
		return 0, err
	}
	rows, err := stmt.Query(item)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err == sql.ErrNoRows {
			fmt.Printf("\nError: %s\nItem: %s does not exist.\n", err, item)
		}
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}
	return id, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateHash(prehash string) []byte {
	hasher := sha256.New()
	data := []byte(prehash)
	hasher.Write(data)
	hashbytes := hasher.Sum(nil)
	return hashbytes
}

func ConvertHashByteSliceToString(hashBytes []byte) string {
	return base64.StdEncoding.EncodeToString(hashBytes)
}

func CompareHash(hash1, hash2 []byte) bool {
	if subtle.ConstantTimeCompare(hash1, hash2) == 1 {
		fmt.Println(true)
		return true
	}
	fmt.Println(false)
	return false
}
