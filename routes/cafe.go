package routes

// authentication required routes for the cafe

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type User struct {
	Id       int
	Username string
	Password string
}

var tokenAuth *jwtauth.JWTAuth

func CafeLoginHandler(w http.ResponseWriter, r *http.Request) error {
	c := CafePageData{
		Page:  "cafeLogin.html",
		Title: "Welcome to the Cafe",
	}
	c.RenderCafe(w)
	return nil
}

func LoginAuthenticator(w http.ResponseWriter, r *http.Request) {
	// okay cool and all but lets be mature about this

	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == "" || password == "" {
		http.Error(w, "Email or password can't be blank.", http.StatusBadRequest)
		return
	}

	user, err := Validate(username, password)
	switch err {
	case 1:
		fmt.Printf("ID: %d", user.Id)
	case 2:
		fmt.Println("User doesnt exist.")
	}

	token := MakeToken(username)

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		// uncomment below for https
		// Secure: true,
		Name:  "jwt", // must be named jwt for jwtauth to properly parse
		Value: token,
	})

	http.Redirect(w, r, "/cafe-portal", http.StatusSeeOther)
}

func MakeToken(name string) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"username": name})
	return tokenString
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1, // this deletes the cookie
		SameSite: http.SameSiteLaxMode,
		Name:     "jwt",
		Value:    "",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Validate(username, password string) (User, int) {
	var u User
	users := []User{
		{Id: 1, Username: "First", Password: "password"},
		{Id: 2, Username: "Second", Password: "secondpassword"},
		{Id: 3, Username: "Third", Password: "yes"},
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return user, 1
		}
	}

	fmt.Println("No user found.")

	return u, 2
}
