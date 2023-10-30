package routes

// 

import (
	"net/http"
	"fmt"
	"log"
	"database/sql"

	"github.com/SmokierLemur51/gleaf/data"
)



func connectDb() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		data.C.Host, data.C.Port, data.C.User, data.C.Password, data.C.DBName
	)
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
	var services []data.Service
	tempUser := data.User{ID: 1,Username: "SmokierLemur51",Password: "password",}
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

	services, err = data.LoadAllServices(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	page := data.PageData{
		Page: "index.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash(tempUser.Username),
		Services: services,
	}
	page.RenderTemplate(w)
}



func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page := data.PageData{
		Page: "about.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	page.RenderTemplate(w)
}


func ContactHandler(w http.ResponseWriter, r *http.Request) {
	page := data.PageData{
		Page: "contact.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	page.RenderTemplate(w)

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	page := data.PageData{
		Page: "register.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	page.RenderTemplate(w)
}

// testing pages
// temp comparison
func ComparisonHandler(w http.ResponseWriter, r *http.Request) {
	tempUser := data.User{
		ID: 1,
		Username: "SmokierLemur51",
		Password: "password",
	}

	page := data.PageData{
		Page: "testing/index_comparison.html",
		Title: "Greenleaf Cleaning",
		Message: MISSION_STATEMENT,
		UserHash: utils.GenerateHash(tempUser.Username),
	}

	page.RenderTemplate(w)

}