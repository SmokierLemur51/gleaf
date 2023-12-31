package handlers

import (
	"net/http"
	"html/template"
	"github.com/SmokierLemur51/gleaf/data"
)

type PublicPageData struct {
	Page     string
	Title    string
	Content  string
	CSS      string
	Services []data.Service
}

var CSS_URL string = "/static/css/testing.css"

func (p PublicPageData) RenderHTMLTemplate(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/" + p.Page)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, p)
	if err != nil {
		return
	}
	// this prevents the superflous hanlder err
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
