package routes

//

import (
	"fmt"
	"net/http"
)

func CafeLoginHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {

	case "GET":
		c := CafePageData{
			Page:  "cafeLogin.html",
			Title: "Welcome to the Cafe",
		}
		c.RenderCafe(w)
		return nil

	case "POST":
		fmt.Printf("Email: %s\tPassword: %s\n\n", r.FormValue("email"), r.FormValue("password"))
	}
	return nil
}
