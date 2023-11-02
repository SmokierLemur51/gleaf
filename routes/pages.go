package routes

import (
	"html/template"
	"net/http"

	"github.com/SmokierLemur51/gleaf/data"
)

type PageData struct {
	Page     string
	Title    string
	UserHash []byte
	Message  string
	Services []data.Service
}

func (p PageData) RenderPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/" + p.Page)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, p)
	if err != nil {
		return
	}
}
