package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// panic(err) needs to be removed and we need to add actual
// error handling in its place



func AddTable(db *sql.DB, tableName string, sqlStatement string) {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\t*    Table '%s' created successfully/\r\n", tableName)
}

func GetValueByColumn(db *sql.DB, returnColumn, tableName, searchColumn, value string) (interface{}, error) {
	// this func is used to return a value (ex the id) of an item in the table
	// by providing the name you know it as ... if that makes sense

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", returnColumn, tableName, searchColumn)
	
	var result interface {}
	
	err := db.QueryRow(query, value).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Value %s not found in table %s", value, table)
		}
		return nil, err
	}
	return result, nil
}

func CreateDevelopmentTables(db *sql.DB) {
	// create service categories table
	createServiceCategoriesTableSQL := `
		CREATE TABLE IF NOT EXISTS service_categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
		);
	`
	// create services table
	createServicesTableSQL := `
		CREATE TABLE IF NOT EXISTS services (
			id SERIAL PRIMARY KEY,
			category_id INTEGER, 
			name TEXT,
			description TEXT,
			cost REAL,
		);
	`
	// create addresses table
	createAddressesTableSQL := `
		CREATE TABLE IF NOT EXISTS addresses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			street TEXT,
			city TEXT,
			state TEXT,
			zip TEXT,
		);
	`
	// create contacts table 
	createContactsTableSQL := `
		CREATE TABLE IF NOT EXISTS contacts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			name TEXT,
			email TEXT,
			phone_number TEXT, 
			FOREIGN KEY (user_id) REFERENCES uesrs(id),
		);
	`
	// create users table
	createUsersTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			pass_hash TEXT,
			contact_id INTEGER,
			address_id INTEGER,
			FOREIGN KEY (contact_id) REFERENCES contacts(id),
			FOREIGN KEY (address_id) REFERENCES addresses(id),
		);
	`
	// create groups table
	createGroupsTableSQL := `
		CREATE TABLE IF NOT EXISTS groups (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
		);
	`
	// create group members table
	createGroupMemberTableSQL := `
		CREATE TABLE IF NOT EXISTS group_members (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			group_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (group_id) REFERENCES groups(id),
		);
	`
	// create bookings table
	createBookingsTableSQL := `
		CREATE TABLE IF NOT EXISTS bookings (
			id INTEGER PRIMARY KEY,
			service_id INTEGER,
			user_id INTEGER,
			address_id INTEGER,
			contact_id INTEGER
			created_at DATE,
			requested_date DATE,
			completed BOOLEAN,
			completed_date DATE,
			cancelled BOOLEAN,
			cancelled_date DATE,
			cancel_id INTEGER,
			paid BOOLEAN,
			payment_id INTEGER,
			group_cleaning BOOL,
			group_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (address_id) REFERENCES addresses(id),
			FOREIGN KEY (contact_id) REFERENCES contacts(id),
			FOREIGN KEY (cancel_id) REFERENCES cancel(id),
			FOREIGN KEY (payment_id) REFERENCES payments(id),
			FOREIGN KEY (group_id) REFERENCES groups(id),
		);
	`
	// create cancelled bookings table
	createCancelledBookingsTableSQL := `
		CREATE TABLE IF NOT EXISTS cancelled_orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_id INTEGER,
			service_id INTEGER,

			FOREIGN KEY (order_id) REFERENCES orders(id),
			FOREIGN KEY (service_id) REFERENCES services(id),

		);
	`
	AddTable(db, "service_categories", createServiceCategoriesTableSQL)
	AddTable(db, "services", createServicesTableSQL)
	AddTable(db, "addresses", createAddressesTableSQL)
	AddTable(db, "contacts", createContactsTableSQL)
	AddTable(db, "users", createUsersTableSQL)
	AddTable(db, "groups", createGroupsTableSQL)
	AddTable(db, "group_members", createGroupMemberTableSQL)
	AddTable(db, "bookings", createBookingsTableSQL)
	AddTable(db, "cancelled_orders", createCancelledBookingsTableSQL)

	fmt.Println("\n\n\n\n\t*    Success creating database tables.")
}

// 
//	Functions to populate to each table () 
//
func InsertServiceCategory(db *sql.DB, name, description string) {
	insertServiceCategorySQL := "INSERT INTO service_categories (name, description) VALUES (?, ?)"
	_, err := db.Exec(insertServiceCategorySQL, name, description)
	if err != nil {
		panic(err)
	}
	fmt.Println("\t*    Successfully inserted into 'services_categories' table\r\n")
}

func InsertService(db *sql.DB, category, name, description string, cost float32) {
	category_id := GetValueByColumn(db, "id", "service_categories", "name", category)
	insertServiceSQL := "INSERT INTO services (category_id, name, description, cost) VALUES (?, ?, ?)"
	_, err := db.Exec(insertServiceSQL, category_id, name, description, cost)
	if err != nil {
		panic(err)
	}
	fmt.Println("\t*    Successfully inserted into 'services' table\r\n")	
}

func InsertAddress(db *sql.DB, street, city, state, zip string) {
	insertAddressSQL := "INSERT INTO addresses (street, city, state, zip) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(insertAddressSQL, street, city, state, zip)
	if err != nil {
		panic(err)
	}
	fmt.Println("\t*    Successfully inserted into 'addresses' table\r\n")	
}

func InsertContact(db *sql.DB, ) {
	user_id := GetValueByColumn(db, "id", "users", "name", name)
	insertContactSQL := "INSERT INTO contacts (user_id, name, email, phone) VALUES (?, ?, ?, ?)"

}

func InsertUser(db *sql.DB) {}
func InsertGroup(db *sql.DB) {}
func InsertGroupMember(db *sql.DB) {}
func InsertBooking(db *sql.DB) {}
func InsertCancelledBooking(db *sql.DB) {}


func PopulateDevelopmentTables(db *sql.DB) {

}