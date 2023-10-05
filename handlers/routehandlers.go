package handlers

import (
	"net/http"
	"fmt"
	"log"
	"database/sql"

	"github.com/SmokierLemur51/gleaf/models"
	"github.com/SmokierLemur51/gleaf/utils"
	"github.com/SmokierLemur51/gleaf/database"
)
 
const (
	MISSION_STATEMENT string = "A cleaner and greener way of living"
	PORT 		= ":5000"
	host 		= "localhost"
	port 		= 5432
	user 		= "postgres"
	password 	= "1lP(=F=<HHwD]v"
	dbname		= "gleaftesting"
)



func connectDb() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	var err error
	var services []models.Service
	tempUser := models.User{ID: 1,Username: "SmokierLemur51",Password: "password",}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	services, err = database.LoadAllServices(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	page := models.PageData{
		Page: "index.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash(tempUser.Username),
		Services: services,
	}

	err = utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "about.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func ContactHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "contact.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "register.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// testing pages
// temp comparison
func ComparisonHandler(w http.ResponseWriter, r *http.Request) {
	tempUser := models.User{
		ID: 1,
		Username: "SmokierLemur51",
		Password: "password",
	}

	page := models.PageData{
		Page: "testing/index_comparison.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash(tempUser.Username),
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}