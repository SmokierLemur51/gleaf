package main

import (
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
  var PORT string = ":5000"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	c := handlers.Controller{}
	c.ConnectDatabase("instance/testing_gorm_v1.db")
	c.RegisterRoutes(r)

	log.Println("Starting server on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
