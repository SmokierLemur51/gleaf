package handlers

import (
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/data"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
)

var (
	key  []byte
	COST = 14
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) ConnectDatabase(databasFile string) {
	var err error
	if c.DB, err = gorm.Open(sqlite.Open(database), databaseFile); err != nil {
		panic(err)
	}
}

func (c Controller) RegisterRoutes(r chi.Router) {
	// test handler
	r.Method(http.MethodGet, "/testing", c.TestHandler())

	// public routes
	r.Method(http.MethodGet, "/", c.IndexHandler())
	r.Method(http.MethodGet, "/about", c.AboutHandler())
	r.Method(http.MethodGet, "/login", c.LoginPageHandler())

	// post methods
	r.Method(http.MethodPost, "/request-estimate/", c.Process_RequestEstimate())
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
		services, err := data.LoadServicesByStatus(c.DB, "active")
		if err != nil {
			log.Printf("Error: %v\r\n", err)
		}

		p := PublicPageData{Page: "index.html", Title: "Greenleaf Cleaning",
			CSS: CSS_URL, Services: services}
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

func (c Controller) Process_RequestEstimate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		est := data.EstimateRequest{
			Name:        r.FormValue("name"),
			Email:       r.FormValue("email"),
			Phone:       r.FormValue("phone"),
			Description: r.FormValue("description"),
			StatusId:    1, // open is the default
		}

		if err := est.InsertEstimateRequest(c.DB); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
