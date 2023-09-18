package database

import (
	"database/sql"
	"fmt"
	"github.com/SmokierLemur51/gleaf/utils"
	_ "github.com/lib/pq"
)

const (
	host 		= "localhost"
	port 		= 5432
	user 		= "postgres"
	password 	= "1lP(=F=<HHwD]v"
	dbname		= "gleaftesting"
)

func InitConn() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	utils.CheckErr(err)

	defer db.Close()

	err = db.Ping()
	utils.CheckErr(err)

	fmt.Println("\tDatabase connection successful!\r\n\n")

	
}


