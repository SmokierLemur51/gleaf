package data
 
import (
    "log"
	"fmt"
    "database/sql"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
    _ "github.com/mattn/go-sqlite3"
)

func CheckExistence(db *sql.DB, table, column, item string) (bool, error) {
    // returns true if it exists
    var count int
    rows, err := db.Query("SELECT COUNT(*) FROM ? WHERE ? = ?", table, column, item)
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

func FindDatabaseID(db *sql.DB, table, column, item string) int {
    rows, err := db.Query(fmt.Sprintf("SELECT id FROM %s WHERE %s = ?", table, column), item)
    if err != nil {
        log.Fatal(err)
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
        log.Fatal(err)
    }
    return id
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