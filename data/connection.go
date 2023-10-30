package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type DBConn struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}


var C = DBConn{
	Host: "localhost", 
	Port: 5432, 
	User: "postgres", 
	Password: "1lP(=F=<HHwD]v",
	DBName: "gleaftesting",
}


func InitConn() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", C.Host, C.Port, C.User, C.Password, C.DBName)
	db, err := sql.Open("postgres", psqlconn)
	CheckErr(err)

	defer db.Close()

	err = db.Ping()
	CheckErr(err)

	fmt.Println("\tDatabase connection successful!\r\n\n")

}


