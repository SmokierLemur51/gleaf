package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SmokierLemur51/gleaf/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	// _ "github.com/lib/pq"
)

const (
	PORT = ":5000"

// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = "1lP(=F=<HHwD]v"
// dbname   = "gleaftesting"
)

// this function is not needed at all ...
// func init() {
// 	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	var db *sql.DB
// 	var err error
// 	db, err = sql.Open("postgres", psqlconn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// }

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// a sample jwt token with claims `user_id:123`
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {

	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// regular routes
	routes.ConfigureRoutes(r)

	// protected routes
	// this one is balogna
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))

		// handle valid & invalid tokens. This example is using provded auth
		// middleware
		r.Use(jwtauth.Authenticator)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("Protected area, hi %v", claims["user_id"])))
		})
	})

	routes.ProtectedRoutes(r)

	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
}
