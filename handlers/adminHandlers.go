package handlers

import (
	"net/http"
	"log"

	"github.com/SmokierLemur51/gleaf/models"
	"github.com/SmokierLemur51/gleaf/database"
	"github.com/SmokierLemur51/gleaf/utils"
)

func AdminIndexHandlers(w http.ResponseWriter, r *http.Request) {
	ServiceCategories, err := database.LoadAllServiceCategories(db)
	if err != nil {
		log.Fatal(err)
	} 
	ActiveServices, err := database.LoadActiveServices(db, ServiceCategories)
	if err != nil {
		log.Fatal(err)
	}
	data := models.PageData{
		Page: "admin/adminIndex.html",
		Title: "CafeGreenleaf",
		UserHash: []byte("HashedPass"),
		Message: "message",
		Services: ActiveServices,
	}
}


// type PageData struct {
// 	Page 		string
// 	Title		string
// 	UserHash 	[]byte
// 	Message 	string
// 	Services    []Service
// }
