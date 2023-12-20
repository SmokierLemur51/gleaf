package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/mattn/go-sqlite3"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
    tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
    _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
    fmt.Printf("\n\nDebug: Sample Token: %s\n\n", tokenString)
}

func main() {
	var PORT string = ":5000"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


    c := handlers.Controller{}
    c.ConnectDatabase("sqlite3", "testing.db")
    c.RegisterRoutes(r)

    

	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
}
