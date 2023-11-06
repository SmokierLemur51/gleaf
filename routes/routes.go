package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
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

	// ** admin pages **
	router.Group(func(r chi.Router) {
		// routes that need auth
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Method(http.MethodGet, "/portal", Handler(CafeLoginHandler))
	})

	router.Group(func(r chi.Router) {
		// post routes for authentication

		r.Post("/portal", LoginAuthenticator)
		r.Post("/logout", LogoutHandler)
	})
}

func Router(tokenAuth *jwtauth.JWTAuth) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Method(http.MethodGet, "/", Handler(IndexHandler))
	r.Method(http.MethodGet, "/about", Handler(AboutHandler))
	r.Method(http.MethodGet, "/contact", Handler(ContactHandler))
	r.Method(http.MethodGet, "/register", Handler(RegisterHandler))

	r.Group(func(r chi.Router) {
		// routes that need auth
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Method(http.MethodGet, "/portal", Handler(CafeLoginHandler))
	})

	r.Group(func(r chi.Router) {
		// post routes for authentication

		r.Post("/portal", LoginAuthenticator)
		r.Post("/logout", LogoutHandler)
	})
	return r
}
