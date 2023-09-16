package main

import (
	"net/http"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/SmokierLemur51/gleaf/database"
	"github.com/SmokierLemur51/gleaf/routes"
)

const (
	PORT = ":5000"
)


func main() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	database.InitConn()
	routes.ConfigureRoutes(r)

	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
}