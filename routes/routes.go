package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned err here
		w.WriteHeader(503)
		w.Write([]byte("Bad"))
	}
}

func ConfigureRoutes(router *chi.Mux) {
	router.Method(http.MethodGet, "/", Handler(IndexHandler))
	router.Method(http.MethodGet, "/about", Handler(AboutHandler))
	router.Method(http.MethodGet, "/contact", Handler(ContactHandler))
	router.Method(http.MethodGet, "/register", Handler(RegisterHandler))

	// testing pages
	router.Method(http.MethodGet, "/comparison", Handler(ComparisonHandler))

	// ** admin pages **
	// router.Get("/admin", handlers.AdminIndexHandler)

}
