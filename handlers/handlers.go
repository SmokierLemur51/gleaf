package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/data"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/go-chi/jwtauth/v5"
)

var (
	key  []byte
	COST = 14
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) ConnectDatabase(database, file string) {
	var err error
	if c.DB, err = sql.Open(database, file); err != nil {
		log.Fatal(err)
	}
	c.DB.Ping()
}

func (c Controller) RegisterRoutes(r chi.Router) {
	// test handler
	r.Method(http.MethodGet, "/controller", c.TestHandler())

	// public routes
	r.Method(http.MethodGet, "/", c.IndexHandler())
	r.Method(http.MethodGet, "/about", c.AboutHandler())
	r.Method(http.MethodGet, "/login", c.LoginPageHandler())

	// post methods
	r.Method(http.MethodPost, "/register-user", c.RegisterNewUser())
}

func (c Controller) Authenticate() { /* pass *jwtauth secret as argument? */ }

func (c Controller) TestHandler() http.HandlerFunc {
	// test authentication here
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{Page: "testing.html", Title: "Success"}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{Page: "index.html", Title: "Greenleaf Cleaning", CSS: CSS_URL, Services: []data.Service{}}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) AboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{Page: "about.html", Title: "Greenleaf Cleaning", CSS: CSS_URL, Services: []data.Service{}}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) LoginPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{Page: "login.html", Title: "Greenleaf Cleaning", CSS: CSS_URL}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) ProcessLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		// hp, err := HashString(r.FormValue("password"), COST)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// u, err := VerifyCredentials(c.DB, r.FormValue("email"), hp)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// sign token with u here
	}
}

func (c Controller) RegisterNewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create new user
	}
}
