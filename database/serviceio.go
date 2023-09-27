package database

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/SmokierLemur51/gleaf/models"
	_ "github.com/lib/pq"
)

// this file is for service table related io


// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *  * * * * * * * 
// * * * * * * * * * * * * * Service Category Related  * * * * * *  * * * * * * * 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *  * * * * * * * 


func InsertServiceCategory(db *sql.DB, name, description string) {
	insertStmt, err := db.Prepare("insert into service_categories (name, description) values ($1, $2);")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(name, description)
	if err != nil {
		log.Fatal(err)
	}
} 

func LoadServiceCategory(db *sql.DB, searchQuery string) (models.ServiceCategory, error) {
	var cat models.ServiceCategory
	
	stmt, err := db.Prepare("select id, name, description from service_categories where name = $1;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	
	err = stmt.QueryRow(searchQuery).Scan(&cat.ID, &cat.Name, &cat.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No matching category found.")
			return cat, err
		} else {
			log.Fatal(err)
		}
	} 

	return cat, nil 
}


func LoadAllServiceCategories(db *sql.DB) ([]models.ServiceCategory, error){
	var ServiceCategoryResults []models.ServiceCategory
	query := "select id, name, description from service_categories;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category models.ServiceCategory
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			fmt.Println("Error scanning row.")
		}
		ServiceCategoryResults = append(ServiceCategoryResults, category)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows: ", err)
	}

	return ServiceCategoryResults, nil
}



// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *  * * * * * * * 
// * * * * * * * * * * * * * * * Service Related  * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *  * * * * * * * 




func InsertService(db *sql.DB, categoryName, name, description string, cost float32, status bool) {
	var category models.ServiceCategory
	var err error
	category, err = LoadServiceCategory(db, categoryName)
	if err != nil {
		return
	}
	insertStmt, err := db.Prepare("insert into services (category_id, name, description, cost, status) values ($1, $2, $3, $4, $5);")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(category.ID, name, description, cost, status)
	if err != nil {
		log.Fatal(err)
	}
}

// func UpdateServicePrice(db *sql.DB, serviceName string, newCost float32) error {
// 	query := "update services set cost = $1 where name = $2;"	

// 	return nil 
// }

func SearchDBForService(db *sql.DB, serviceName string) (models.Service, error) {
	var service models.Service 
	query := "select id, category_id, name, description, cost from services where name = $1;"
	rows, err := db.QueryRow(query, serviceName)
	if err != nil {
		return service, err
	}
	defer rows.Close()


	return service, nil  	
}


func AlterServiceStatus(db *sql.DB, unalteredService, alteredService models.Service) error {
	var service models.Service
	query := 

	return service, nil
}



func LoadActiveServices(db *sql.DB, serviceCategories []models.ServiceCategory) ([]models.Service, error) {
	var ActiveServices []models.Service
	query := "select * from services where status = true"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Service
		if item.Status == true {
			err := rows.Scan(&item.ID, &item.Type_ID, &item.Name, &item.Description, &item.Cost, &item.Status)
			if err != nil {
				fmt.Println("Error scanning row: ", err)
			}
			for _, category := range serviceCategories {
				if item.Type_ID == category.ID {
					item.CategoryName = category.Name
					break
				}
			}
			ActiveServices = append(ActiveServices, item)
		}
	}

	return ActiveServices, nil
}


func ChangeServiceStatus(db *sql.DB, serviceName string, status bool) {
	return
}

// type Service struct {
// 	ID			int16
// 	Type_ID		int16
// 	Type 		string
// 	Description string
// 	Cost 		float32
// }

// 	createServicesTableSQL := `
// 		CREATE TABLE IF NOT EXISTS services (
// 			id SERIAL PRIMARY KEY,
// 			category_id INTEGER, 
// 			name TEXT,
// 			description TEXT,
// 			cost REAL,
// 		);