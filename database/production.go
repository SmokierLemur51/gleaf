package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func TestPostgresConn(db *sql.DB, USERNAME string, DBNAME string) {
	connectionString := "user="+USERNAME+" "+"dbname="+DBNAME+" sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successful")
}