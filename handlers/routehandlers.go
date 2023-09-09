package handlers

import (
	"net/http"

	"github.com/SmokierLemur51/gleaf/models"
	"github.com/SmokierLemur51/gleaf/utils"
)

const (
	MISSION_STATEMENT string = "A cleaner and greener way of living"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tempUser := models.User{
		ID: 1,
		Username: "SmokierLemur51",
		Password: "password",
	}

	page := models.PageData{
		Page: "index.html",
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

func ContactHandler(w http.ResponseWriter, r *http.Request) {
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