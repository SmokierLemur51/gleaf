package data

import (
	"database/sql"
	"log"
	"fmt"
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

func LoadServiceCategory(db *sql.DB, searchQuery string) (ServiceCategory, error) {
	var cat ServiceCategory
	
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



func LoadAllServiceCategories(db *sql.DB) ([]ServiceCategory, error){
	var ServiceCategoryResults []ServiceCategory
	query := "select id, name, description from service_categories;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category ServiceCategory
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			fmt.Println("Error scanning row.")
		}
		ServiceCategoryResults = append(ServiceCategoryResults, category)
	}

	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No services found.")
		}
		fmt.Println("Error iterating over rows: ", err)
	}

	return ServiceCategoryResults, nil
}



// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *  * * * * * * * 
// * * * * * * * * * * * * * * * Service Related  * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *  * * * * * * * 




func InsertService(db *sql.DB, categoryName, name, description string, cost float32, status bool) (error) {
	var category ServiceCategory
	var err error
	category, err = LoadServiceCategory(db, categoryName)
	if err != nil {
		return err
	}
	insertStmt, err := db.Prepare("insert into services (category_id, name, description, cost, status) values ($1, $2, $3, $4, $5);")
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(category.ID, name, description, cost, status)
	if err != nil {
		return err
	}
	return nil 	
}

// func UpdateServicePrice(db *sql.DB, serviceName string, newCost float32) error {
// 	query := "update services set cost = $1 where name = $2;"	

// 	return nil 
// }

func SearchDBForService(db *sql.DB, serviceName string) (Service, error) {
	var service Service 
// 	query := "select id, category_id, name, description, cost from services where name = $1;"
// 	rows, err := db.QueryRow(query, serviceName)
// 	if err != nil {
// 		return service, err
// 	}
// 	defer rows.Close()


	return service, nil
}


func AlterServiceStatus(db *sql.DB, unalteredService, alteredService Service) (Service, error) {
	var service Service
	// query := 

	return service, nil
}

func LoadAllServices(db *sql.DB) ([]Service, error) {
	stmt := "select id, category_id, name, description, cost, status from services;"
	rows, err := db.Query(stmt)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer rows.Close()

	var services []Service
	for rows.Next() {
		var service Service 
		if err := rows.Scan(&service.ID, &service.Type_ID, &service.Name, &service.Description, &service.Cost, &service.Status); err != nil {
			fmt.Printf("Error: %s", err)
		}
		services = append(services, service)
	}
	// check for err in rows 
	if err := rows.Err(); err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, service := range services {
		fmt.Printf("\n\tName: %s\n", service.Name)
	}
	return services, nil 
}  

// type Service struct {
// 	ID			 int 	   	`db:"id"`
// 	Type_ID		 int 		`db:"category_id"`
// 	CategoryName string
// 	Name 	  	 string  	`db:"name"`
// 	Description  string		`db:"description"`
// 	Cost 		 float32 	`db:"cost"`
// 	Status       bool 		`db:"status"`
// }


func LoadActiveServices(db *sql.DB, serviceCategories []ServiceCategory) ([]Service, error) {
	var ActiveServices []Service
	query := "select * from services where status = true"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item Service
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