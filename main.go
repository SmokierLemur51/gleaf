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


// this function is not needed at all ...
func init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// description := "A deep cleanse for your old home. It will seem as if you never even lived there..."
	// err = database.InsertService(db, "moving", "Move Out Deep Clean", description, 200.00, true)
	// if err != nil {
	// 		log.Println(err)
	// }
	// helpingHand := "Falling behind on cleaning? Don't worry, we've got your back."
	// err = database.InsertService(db, "residential", "Helping Hand", helpingHand, 125.00, true)
	// if err != nil {
	// 	log.Println(err)
	// }
}


func main() {

	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(r)
	

	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
}