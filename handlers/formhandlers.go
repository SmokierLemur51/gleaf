package handlers

// import (
// 	"net/http"
// 	"log"

// 	"github.com/SmokierLemur51/gleaf/models"
// )
//  - - - - - - - -  - - - - - - - -  - - - - -
//  - - - - - - - - POST HANDLERS - - - - - - -
//  - - - - - - - -  - - - - - - - -  - - - - -  
// func HandleScheduleInput(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// name := r.FormValue("name")
	// email := r.FormValue("email")
	// date := r.FormValue("date")
	// redirect to homepage ... 
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }


// func ProcessContactInfo(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	contact := models.ContactInfo {
// 		Name: r.FormValue("name"),
// 		Email: r.FormValue("email"),
// 		Number: r.FormValue("number"),
// 	}
	// form validation
// 	if contact.Email == "" || contact.Name == "" {
// 		http.Error(w, "Name and email are required fields.", http.StatusBadRequest)
// 		return
// 	}
// 	log.Printf("Recieved this informaiton: %+v", contact)
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }