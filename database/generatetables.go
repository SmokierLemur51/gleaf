package database

import (
	"database/sql"
	"fmt"
	"time"
    "log"

    "github.com/SmokierLemur51/gleaf/utils"
	_ "github.com/mattn/go-sqlite3"
)

// panic(err) needs to be removed and we need to add actual
// error handling in its place


func CreateGleafTables(db *sql.DB) {
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
            tenant_name TEXT
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
			phone_number TEXT,
			email TEXT,
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
            creator INTEGER,
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
    // create completed bookings CreateGleafTables
    createCompletedBookingsTableSQL := `
        CREATE TABLE IF NOT EXISTS completed_orders (
            id INTEGER PRIMARY KEY AUTOINCREMENT,    
            order_id INTEGER,
            service_id INTEGER,
            booking_date DATE,
            completed_date DATE,
            booking_price REAL,
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
	AddTable(db, "completed_orders", createCancelledBookingsTableSQL)

	fmt.Println("\n\n\n\n\t*    Success creating database tables.")
}



// 
//	Functions to populate to each table () 
//		These should be updated to return errors to prevent crashing
//

func InsertAddress(db *sql.DB, name, street, city, state, zip string) {
	insertAddressSQL := "INSERT INTO addresses (tenant_name, street, city, state, zip) VALUES (?, ?, ?, ?, ?);"
	_, err := db.Exec(insertAddressSQL, name, street, city, state, zip)
	if err != nil {
		panic(err)
	}
	fmt.Println("\t*    Successfully inserted into 'addresses' table\r\n")	
}


func InsertContact(db *sql.DB, username, name, email, phone string) {
	// if username is not empty, find user_id
	if username != "" || username != nil {
		user_id, err := GetValueByColumn(db, "id", "contacts", "name", name)
		if err != nil {
			panic(err)
		}
	} else {
		// if username is empty, user_id is null because og does not exist
		user_id = "NULL"
	}
	insertContactSQL := "INSERT INTO contacts (user_id, name, email, phone) VALUES (?, ?, ?, ?);"
	_, err := db.Exec(insertContactSQL, user_id, name, email, phone)
		if err != nil {
			panic(err)
		}
		fmt.Println("\t*    Successfully inserted into 'contacts' table\r\n")	
}



func InsertUser(db *sql.DB, username, password, name, email, phone, street, city, state, zip string) {
	contact_id, err := GetContactID(db, phone, email)
    if err != nil {
        log.Panic(err)
    }
    switch contact_id{
        case -1:
            // contact does not exist
            InsertContact(db, "NULL", name, email, phone)
        case -2:
            // FIXLOG 2
            fmt.Println("Contact exists but does not match.")
    }
    address_id, err := GetAddressID(db, name, street, city, state, zip)
    if err != nil {
        log.Panic(err)
        return
    }
    switch address_id{
        case -1:
            InsertAddress(db, name, street, city, state, zip)
            address_id, err = GetAddressID(db, name, street, city, state, zip)
            if err != nil {
                log.Panic(err)
            }
        case -2:
            // FIXLOG 3
            fmt.Println("Address exists but does not match ...")
    }
    if (address_id > 0 && contact_id > 0) {
    	// FIXLOG 4 
        hashed_pass := utils.ConvertHashByteSliceToString(utils.GenerateHash(password))
    	insertUserSQL := "INSERT INTO users (username, pass_hash, contact_id, phone_number, email, address_id) VALUES (?, ?, ?, ?, ?, ?);"
    	_, err := db.Exec(insertUserSQL, username, hashed_pass, contact_id, email, phone, address_id)
    } else {
        fmt.Println("Cannot create user due to '%s'. Temporary Exit", err)
        return 
    }
}


func InsertGroup(db *sql.DB, groupName, creatorUsername string) {
    unique, err := DoesGroupExist(db, groupName, creatorUsername)
    if err != nil {
        log.Errorf(err)
    }
    switch unique {
        case -1:
            fmt.Println("Group Name cannot be empty.")
        case -2:
            fmt.Println("Must be signed in to create group.")
        case 1:
            insertGroupSQL := "INSERT INTO groups (name, creator) VALUES (?, ?);"
            _, err := db.Exec(insertGroupSQL, groupName, creatorUsername)
            if err != nil {
                log.Errorf(err)
            }
    }    
}


func InsertGroupMember(db *sql.DB) {

}


func InsertBooking(db *sql.DB, createDate, requestedDate, cancelDate, completeDate time.Time, userInformation models.CurrentUser,serviceID, 
										groupID int, paymentInfo models.PaymentInformation,) {
	
}

	// createBookingsTableSQL := `
	// 	CREATE TABLE IF NOT EXISTS bookings (
	// 		id INTEGER PRIMARY KEY,
	// 		service_id INTEGER,
	// 		user_id INTEGER,
	// 		address_id INTEGER,
	// 		contact_id INTEGER
	// 		created_at DATE,
	// 		requested_date DATE,
	// 		completed BOOLEAN,
	// 		completed_date DATE,
	// 		cancelled BOOLEAN,
	// 		cancelled_date DATE,
	// 		cancel_id INTEGER,
	// 		paid BOOLEAN,
	// 		payment_id INTEGER,
	// 		group_cleaning BOOL,
	// 		group_id INTEGER,
	// 		FOREIGN KEY (user_id) REFERENCES users(id),
	// 		FOREIGN KEY (address_id) REFERENCES addresses(id),
	// 		FOREIGN KEY (contact_id) REFERENCES contacts(id),
	// 		FOREIGN KEY (cancel_id) REFERENCES cancel(id),
	// 		FOREIGN KEY (payment_id) REFERENCES payments(id),
	// 		FOREIGN KEY (group_id) REFERENCES groups(id),
	// 	);
	// `


func InsertCancelledBooking(db *sql.DB) {}


func PopulateDevelopmentTables(db *sql.DB) {

}