package routes

//

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/data"
	"github.com/SmokierLemur51/gleaf/utils"
)

const (
	MISSION_STATEMENT = "Mission Statement"
)

func connectDb() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", data.C.Host, data.C.Port, data.C.User, data.C.Password, data.C.DBName)
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

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	var db *sql.DB
	var err error
	var services []data.Service
	tempUser := data.User{ID: 1, Username: "SmokierLemur51", Password: "password"}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", data.C.Host, data.C.Port, data.C.User, data.C.Password, data.C.DBName)
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

	page := PageData{
		Page:     "index.html",
		Title:    "Greenleaf Cleaning",
		Message:  MISSION_STATEMENT,
		UserHash: utils.GenerateHash(tempUser.Username),
		Services: services,
	}
	page.RenderPage(w)
	return nil
}

func AboutHandler(w http.ResponseWriter, r *http.Request) error {
	page := PageData{
		Page:     "about.html",
		Title:    "Greenleaf Cleaning",
		Message:  MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	page.RenderPage(w)
	return nil
}

func ContactHandler(w http.ResponseWriter, r *http.Request) error {
	page := PageData{
		Page:     "contact.html",
		Title:    "Greenleaf Cleaning",
		Message:  MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	page.RenderPage(w)
	return nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) error {
	page := PageData{
		Page:     "register.html",
		Title:    "Greenleaf Cleaning",
		Message:  MISSION_STATEMENT,
		UserHash: utils.GenerateHash("test"),
	}
	page.RenderPage(w)
	return nil
}

// testing pages
// temp comparison
func ComparisonHandler(w http.ResponseWriter, r *http.Request) error {
	tempUser := data.User{
		ID:       1,
		Username: "SmokierLemur51",
		Password: "password",
	}

	page := PageData{
		Page:     "testing/index_comparison.html",
		Title:    "Greenleaf Cleaning",
		Message:  MISSION_STATEMENT,
		UserHash: utils.GenerateHash(tempUser.Username),
	}

	page.RenderPage(w)
	return nil
}
