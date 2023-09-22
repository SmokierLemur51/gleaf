package routes

import (
	"github.com/SmokierLemur51/gleaf/handlers"

	"github.com/go-chi/chi/v5"
)


func ConfigureRoutes(router *chi.Mux) {
	router.Get("/", handlers.IndexHandler)
	router.Get("/about", handlers.AboutHandler)
	router.Get("/contact", handlers.ContactHandler)
	router.Get("/register", handlers.RegisterHandler)
	// testing pages
	router.Get("/comparison", handlers.ComparisonHandler)

	// ** admin pages **
	router.Get("/admin", handlers.AdminIndexHandler)

}

