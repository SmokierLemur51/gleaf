package routes

import (

	"github.com/go-chi/chi/v5"
)


func ConfigureRoutes(router *chi.Mux) {
	router.Get("/", IndexHandler)
	router.Get("/about-us", AboutHandler)
}

