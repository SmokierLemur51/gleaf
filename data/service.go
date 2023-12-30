/*
Things you can get from service.go

Service Type

	Methods
	- InsertService(db) -> Inserts a new service after calling CheckExistingService()
	- UpdateService(db) -> Updates service
	- RunReport(db, period) -> With period being day, week etc. Generates reports from that service

Functions
- CheckExistingService(db, service) -> Returns true if exists, false if not
- CheckServiceFields(*Service) -> Check that required fields exist
*/
package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// status 1 active, status 2 inactive
type Service struct {
	Id          int     `json:"id" db:"id"`
	CategoryId  int     `json:"category_id" db:"category_id"`
	Status      int     `json:"status_id" db:"status_id"`
	Service     string  `json:"service" db:"service_name"`
	Description string  `json:"description" db:"service_description"`
	Selling     float64 `json:"selling" db:"selling"`
	Category    ServiceCategory
}

func CheckServiceFields(s *Service) (*Service, error) { return s, nil }

func (s Service) InsertService(db *sql.DB) error {
	var execute bool
	var err error
	execute, err = CheckExistence(db, "services", "service", s.Service)
	if err != nil {
		log.Println(err)
		return err
	}
	// remember, the check existing returns true if the product already exists, so it skips
	switch execute {
	case false:
		_, err := db.Exec(
			"INSERT INTO services (category_id, status_id, service_name, service_description, selling) VALUES (?,?,?,?,?)",
			s.CategoryId, s.Status, s.Service, s.Description, s.Selling,
		)
		if err != nil {
			return err
		}
	case true:
		fmt.Printf("Service %s already exists.\n", s.Service)
	}
	return nil
}

func (s *Service) UpdateSellingPrice(db *sql.DB, newSelling float64) error { return nil }

func (s *Service) LoadServiceByName(db *sql.DB, services string) error {
	query := "SELECT id, category_id, status_id, service, description, selling FROM services WHERE service = ?"
	if err := db.QueryRow(query).Scan(&s.Id, &s.CategoryId, &s.Status, &s.Service, &s.Description, &s.Selling); err != nil {
		log.Printf("Err: %v", err)
		return err
	}
	return nil
}

// converting this to the method above
func LoadServiceByName(db *sql.DB, service string) (Service, error) {
	rows, err := db.Query(
		"SELECT id, category_id, status_id, service, description, selling FROM services WHERE service = ?",
		service,
	)
	if err != nil {
		log.Fatal(err)
	}
	var s Service
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.CategoryId, &s.Status, &s.Service, &s.Description, &s.Selling)
		if err != nil {
			log.Fatal(err)
			return s, err
		}
	}
	return s, nil
}

func LoadServicesByStatus(db *sql.DB, status string) ([]Service, error) {
	var services []Service
	rows, err := db.Query(
		"SELECT id, category_id, status_id, service, description, selling FROM services WHERE status_id = ?",
		status,
	)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var s Service
		err := rows.Scan(&s.Id, &s.CategoryId, &s.Status, &s.Service, &s.Description, &s.Selling)
		if err != nil {
			log.Fatal(err)
		}
		services = append(services, s)
		log.Printf("Added service: %s to %s slice.", s.Service, status)
	}
	return services, nil
}

// func LoadServicesByCategory(db *sql.DB, category string) ([]Service, error) {
//     var services []Service
//     rows, err := db.Query(
//         "SELECT id, category_id, status_id, service, description, selling FROM services WHERE category_id = ?",
//         FindCategoryId(db, category),
//     )
//     if err != nil {
//         log.Fatal(err)
//     }
//     for rows.Next() {
//         var s Service
//         err := rows.Scan(&s.Id, &s.CategoryId, &s.Status, &s.Service, &s.Description, &s.Selling)
//         if err != nil {
//             log.Fatal(err)
//         }
//         services = append(services, s)
//         log.Printf("Added service: %s to %s slice.", s.Service, category)
//     }
//     return services, nil
// }

func PopulateServicesTable(db *sql.DB, servs []Service) error {
	for _, p := range servs {
		if err := p.InsertService(db); err != nil {
			log.Printf("Error with service %s:\t%v", p.Service, err)
		}
	}

	return nil
}

func EarlyStageServiceSlice() []Service {

	return []Service{
		// moving
		{Service: "Move In/Out Deep Cleanse", CategoryId: 1, Status: 1,
			Description: "A deep cleanse of your home before or after you move out.", Selling: 275.00},
		// residential
		{Service: "Quick Clean", CategoryId: 2, Status: 1,
			Description: "A quick cleaning of your house.", Selling: 150.00},
		{Service: "Residential Deep Clean.", CategoryId: 2, Status: 1,
			Description: "Deep clean of carpet, furniture, fridge and more.", Selling: 500.00},
		{Service: "Carpet Cleaning", CategoryId: 2, Status: 5,
			Description: "Remove stains, pet smells, and allergens from your carpet. A deep cleanse and shampooing.", Selling: 250.00},
		// commercial
			{Service: "Office Cleaning", CategoryId: 6, Status: 1,
			Description: "Make your office feel more like home. Stop wasting time cleaning it yourself!", Selling: 300.00},
		// residential exterior 3
		{Service: "Gutter Cleaning", CategoryId: 3, Description: "Remove bird nests, leaves and sticks, or anything else that hinders the flow of water.", Selling: 300.00},
		{Service: "Window Cleaning",CategoryId: ,Status: ,Description: , Selling: ,},
	}
}
