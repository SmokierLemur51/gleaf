package main

import (
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
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
	c.ConnectDatabase("sqlite3", "instance/testing.db")
	c.RegisterRoutes(r)

	u := handlers.User{Email: "ldl6147@gmail.com"}
	u.InsertUser(c.DB)

	log.Println("Starting server on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
