package data

import (
	"database/sql"
	"fmt"
	// "time"
    "log"

    // "github.com/SmokierLemur51/gleaf/utils"
	_ "github.com/lib/pq"
)

// panic(err) needs to be removed and we need to add actual
// error handling in its place

func AddTable(db *sql.DB, tableName, sqlStatement string) error {

	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	fmt.Printf("\t*    Table '%s' created successfully\r\n", tableName)
	return err
}





func CreateGleafTables(db *sql.DB) {
	// create service categories table
	// createServiceCategoriesTableSQL := `
	// 	CREATE TABLE IF NOT EXISTS service_categories (
	// 		id SERIAL PRIMARY KEY,
	// 		name VARCHAR () NOT NULL,
	// 		description VARCHAR () NOT NULL,
	// 	);`

	// create services table
	createServicesTableSQL := `
		CREATE TABLE IF NOT EXISTS services (
			id SERIAL PRIMARY KEY,
			category_id INTEGER, 
			name VARCHAR (50),
			description VARCHAR (250),
			cost REAL
		);
	`
	// create addresses table
	createAddressesTableSQL := `
		CREATE TABLE IF NOT EXISTS addresses (
			id SERIAL PRIMARY KEY,
            tenant_name VARCHAR (50),
			street VARCHAR (50) NOT NULL,
			city VARCHAR (25) NOT NULL,
			state VARCHAR (2) NOT NULL,
			zip VARCHAR (5) NOT NULL
		);
	`
	// create contacts table 
	createContactsTableSQL := `
		CREATE TABLE IF NOT EXISTS contacts (
			id SERIAL PRIMARY KEY,
			name VARCHAR (50) NOT NULL,
			email VARCHAR (50) NOT NULL,
			phone_number VARCHAR (10) NOT NULL
		);
	`
	// create users table
	createUsersTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR (50) NOT NULL,
			pass_hash VARCHAR (60) NOT NULL,
			contact_id INTEGER REFERENCES contacts(id),
			phone_number VARCHAR (10),
			email VARCHAR (50) NOT NULL,
			address_id INTEGER REFERENCES addresses(id)
		);
	`
	// create groups table
	createGroupsTableSQL := `
		CREATE TABLE IF NOT EXISTS groups (
			id SERIAL PRIMARY KEY,
			name VARCHAR (50) NOT NULL,
            creator INTEGER REFERENCES users(id)
		);
	`
	// create group members table
	createGroupMemberTableSQL := `
		CREATE TABLE IF NOT EXISTS group_members (
			id SERIAL PRIMARY KEY,
			group_id INTEGER REFERENCES groups(id),
			user_id INTEGER REFERENCES users(id)
		);
	`
	// create bookings table
	createBookingsTableSQL := `
		CREATE TABLE IF NOT EXISTS bookings (
			id INTEGER PRIMARY KEY,
			service_id INTEGER REFERENCES services(id),
			user_id INTEGER REFERENCES users(id),
			address_id INTEGER REFERENCES addresses(id),
			contact_id INTEGER REFERENCES contacts(id),
			created_at DATE,
			requested_date DATE,
			completed BOOLEAN,
			completed_date DATE,
			cancelled BOOLEAN,
			cancelled_date DATE,
			paid BOOLEAN,
			group_cleaning BOOL,
			group_id INTEGER REFERENCES groups(id)
		);
	`
	// create cancelled bookings table
	createCancelledBookingsTableSQL := `
		CREATE TABLE IF NOT EXISTS cancelled_bookings (
			id SERIAL PRIMARY KEY,
			order_id INTEGER REFERENCES bookings(id),
			service_id INTEGER REFERENCES services(id)
		);
	`
    // create completed bookings CreateGleafTables
    createCompletedBookingsTableSQL := `
        CREATE TABLE IF NOT EXISTS completed_bookings (
            id SERIAL PRIMARY KEY,    
            order_id INTEGER REFERENCES bookings(id),
            service_id INTEGER REFERENCES services(id),
            booking_date DATE,
            completed_date DATE,
            booking_price REAL
        );
    `
	// AddTable(db, "service_categories", createServiceCategoriesTableSQL)
	if err := AddTable(db, "services", createServicesTableSQL); err != nil{
		log.Fatal(err)
	}
	AddTable(db, "addresses", createAddressesTableSQL)
	AddTable(db, "contacts", createContactsTableSQL)
	AddTable(db, "users", createUsersTableSQL)
	AddTable(db, "groups", createGroupsTableSQL)
	AddTable(db, "group_members", createGroupMemberTableSQL)
	AddTable(db, "bookings", createBookingsTableSQL)
	AddTable(db, "cancelled_bookings", createCancelledBookingsTableSQL)
	AddTable(db, "completed_bookings", createCompletedBookingsTableSQL)

	fmt.Println("\n\n\n\n\t*    Success creating database tables.\r\n\n")
	db.Close()
}