package admin

import (
	"html/template"
	"net/http"

	"github.com/SmokierLemur51/gleaf/data"
)

type AdminPageData struct {
	Page                      string
	Title                     string
	ServiceCategories         []data.ServiceCategory
	Services                  []data.Service
	IncompleteContactRequests []data.ContactRequest
	IncompleteBookings        []data.Bookings
	FinancialData             data.Finances
}

func (p AdminPageData) RenderAdminPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/" + p.Page)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, p)
	if err != nil {
		return
	}
}
