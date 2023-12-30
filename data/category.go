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

// add error return
func (s *ServiceCategory) InsertCategory(db *sql.DB) {
	var execute bool
	var err error
	execute, err = CheckExistence(db, "service_categories", "category", s.Category)
	if err != nil {
		log.Println(err)
		return
	}
	// if true, it exists and therefore no insertion
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

func UpdateCategoryStruct(cat *ServiceCategory, newCategory, newAdminInfo, newPublicInfo string) *ServiceCategory {
	cat.Category = newCategory
	cat.AdminInformation = newAdminInfo
	cat.PublicInformation = newPublicInfo
	return cat
}

func (s ServiceCategory) UpdateServiceCategory(db *sql.DB) error {
	// check and make sure the original struct already
	stmt, err := db.Prepare("UPDATE service_categories SET category = ?, admin_information = ?, public_information = ? WHERE id = ?")
	if err != nil {
		log.Printf("Error '%v' in UpdateServiceCategory method, db.Prepare stmt. \n", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Category, s.AdminInformation, s.PublicInformation, s.ID)
	if err != nil {
		log.Printf("Error '%v' in UpdateServiceCat method, stmt.Exec.", err)
		return err
	}

	log.Printf("Successfully updated Category: %s", s.Category)
	return nil
}

// end methods
func LoadAllCategories(db *sql.DB) ([]ServiceCategory, error) {
	var cats []ServiceCategory

	rows, err := db.Query("SELECT * FROM service_categories")
	if err != nil {
		log.Printf("Error: '%v' in 'LoadAllCategories()' query\n", err)
		return []ServiceCategory{}, err
	}
	defer rows.Close()
	// iterate anc check for nil/NULL rows
	for rows.Next() {
		if err == sql.ErrNoRows {
			// i dont think this is even in the correct spot
			log.Println("No categories found during query")
			return []ServiceCategory{}, err
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

func LoadCategory(db *sql.DB, category string) (ServiceCategory, error) {
	var c ServiceCategory
	id, err := FindDatabaseID(db, "service_categories", "category", category)
	if err != nil {
		return c, err
	}
	// next two lines are part of one db.QueryRow statement, wanted to spread it out some
	if err := db.QueryRow("SELECT id, category, admin_information, public_information FROM service_categories WHERE id = ?",
		id).Scan(&c.ID, &c.Category, &c.AdminInformation, &c.PublicInformation); err != nil {
		return ServiceCategory{}, err
	}
	// return the single category
	return c, nil
}
