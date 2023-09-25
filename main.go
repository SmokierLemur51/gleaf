package main

import (
	"net/http"
	"log"
	"fmt"
	"database/sql"

	"github.com/go-chi/chi/v5"
	// "github.com/SmokierLemur51/gleaf/database"
	"github.com/SmokierLemur51/gleaf/routes"
	_ "github.com/lib/pq"
)

const (
	PORT = ":5000"
	host 		= "localhost"
	port 		= 5432
	user 		= "postgres"
	password 	= "1lP(=F=<HHwD]v"
	dbname		= "gleaftesting"
)

var db *sql.DB


func init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}


func main() {

	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(r)
	

	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
	defer db.Close()
}