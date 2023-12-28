/*
Things you can get from category.go

Category Type

	Methods
	- InsertCategory(db) -> Inserts a new category after calling CheckExistingCategory()
	- UpdateCategory(db) -> Updates category
	- RunReport(db, period) -> With period being day, week etc. Generates reports from that category

Functions
  - LoadAllCategories(db) -> Load all categories of course my dear
*/
package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ServiceCategory struct {
	ID                int    `json:"id" db:"id"`
	Category          string `json:"category" db:"category"`
	AdminInformation  string `json:"adminInformation" db:"admin_information"`
	PublicInformation string `json:"publicInformation" db:"public_information"`
}

func (s *ServiceCategory) InsertCategory(db *sql.DB) {
	var execute bool
	var err error
	execute, err = CheckExistence(db, "service_categories", "category", s.Category)
	if err != nil {
		log.Println(err)
		return
	}
	// remember, the check existing returns true if the product already exists, so it skips
	switch execute {
	case false:
		_, err := db.Exec(
			"INSERT INTO service_categories (category, admin_information, public_information) VALUES (?,?,?)",
			s.Category, s.AdminInformation, s.PublicInformation,
		)
		if err != nil {
			log.Fatal(err)
		}
	case true:
		fmt.Printf("Category %s already exists.\n", s.Category)
	}
}

func (s *ServiceCategory) UpdateServiceCategory(db *sql.DB, category, adminInfo, publicInfo string) error {
	// check and make sure the original struct already

	return nil
}

// end methods
func LoadAllCategories(db *sql.DB) ([]ServiceCategory, error) {
	var cats []ServiceCategory

	rows, err := db.Query("SELECT * FROM service_categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// iterate anc check for nil/NULL rows
	for rows.Next() {
		if err == sql.ErrNoRows {
			// i dont think this is even in the correct spot
			log.Println("No categories found during query")
			return cats, nil
		}
		var c ServiceCategory
		err := rows.Scan(&c.ID, &c.Category, &c.AdminInformation, &c.PublicInformation)
		if err != nil {
			log.Println(err)
		}
		cats = append(cats, c)
	}

	return cats, nil
}

func LoadCategory(db *sql.DB) (ServiceCategory, error) { var c ServiceCategory; return c, nil }
