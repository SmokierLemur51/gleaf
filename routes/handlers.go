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
	move              = "A deep clean of your old space, make it look like you were never there."
	resDeep           = "A deep cleaning of your house to help you catch up on what you were behind on."
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

func tempServiceLoader() ([]data.Service, error) {
	return []data.Service{
		{ID: 1, Type_ID: 1, CategoryName: "Residential", Name: "Quick Clean", Description: "A quick clean of your house!", Cost: 150.00, ImageURL: "/static/img/quick-clean-smaller.jpg"},
		{ID: 2, Type_ID: 2, CategoryName: "Window", Name: "Eco-Friendly Window Cleaning", Description: "Window deep cleaning", Cost: 20.00, ImageURL: "/static/img/eco-window.jpeg"},
		{ID: 3, Type_ID: 3, CategoryName: "Moving", Name: "Move Out Cleaning", Description: move, Cost: 350.00, ImageURL: "/static/img/moving.png"},
		{ID: 4, Type_ID: 1, CategoryName: "Residential", Name: "Residential Deep Clean", Description: resDeep, Cost: 300.00, ImageURL: "/static/img/res-clean.jpg"},
	}, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	// the below is edited out until postgres ise

	// var db *sql.DB
	// var err error
	// var services []data.Service
	// tempUser := data.User{ID: 1, Username: "SmokierLemur51", Password: "password"}
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", data.C.Host, data.C.Port, data.C.User, data.C.Password, data.C.DBName)
	// db, err = sql.Open("postgres", psqlconn)
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer db.Close()

	// services, err = data.LoadAllServices(db)

	services, err := tempServiceLoader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	page := PageData{
		Page:     "index.html",
		Title:    "Greenleaf Cleaning",
		Message:  MISSION_STATEMENT,
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
