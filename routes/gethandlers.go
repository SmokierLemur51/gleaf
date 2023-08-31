package routes

import (
	"net/http"

	"github.com/SmokierLemur51/gleaf/models"
	"github.com/SmokierLemur51/gleaf/utils"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "index.html",
		Title: "Greenleaf Cleaning",
		Message: "Fucking hell its all gone, time to start over ... ",
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "about.html",
		Title: "About Us",
		Message: "We are an industry leading company...",
	}
	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

