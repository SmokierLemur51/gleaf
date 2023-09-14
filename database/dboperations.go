package data

import (
    "fmt"
    "log"
    "database/sql"
    // "main/models"

    _ "github.com/mattn/go-sqlite3"
)

func AddTable(db *sql.DB, tableName string, sqlStatement string) {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\t*    Table '%s' created successfully/\r\n", tableName)
}

// FIXLOG 5
func GetValueByColumn(db *sql.DB, returnColumn, tableName, searchColumn, value string) (interface{}, error) {
	// this func is used to return a value (ex the id) of an item in the table
	// by providing the name you know it as ... if that makes sense
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?;", returnColumn, tableName, searchColumn)
	var result interface {}
	err := db.QueryRow(query, value).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Value %s not found in table %s", value, tableName)
		}
		return nil, err
	}
	return result, nil
}


func GetIDFromTableColValue(db *sql.DB, tableName, searchColumn, searchValue string) (int, error) {
	var result int
	query := fmt.Sprintf("SELECT id FROM %s WHERE %s = ?;", tableName, searchColumn)
	err := db.QueryRow(query, searchValue).Scan(&result)
	if err == sql.ErrNoRows {
		// does not exist
		return -1, err
	} else if err != nil {
		// error executing query
		return -1, err
	}
	return result, nil
}

func UsernameAvailabilityChecker(db *sql.DB, username string) (bool, error) {
    // return true if available, false if taken
    var status bool
    var value int
    query := "SELECT count(*) FROM users WHERE username = ?;"
    err := db.QueryRow(query, username).Scan(&value)
    if err != nil {
        log.Panic(err)
    }
    if value == 0 {
        status = true
    } else if value > 0 {
        fmt.Printf("Username %s already taken.", username)
        status = false
    }
    return status, err 
}

func GetContactID(db *sql.DB, phone, email string) (int, error) {
	// this one could definitely be used to improve the other
	query := fmt.Sprintf("SELECT id FROM contacts WHERE phone_number = ? AND email = ? LIMIT 1;")
	var contact_id int
	err := db.QueryRow(query, phone, email).Scan(&contact_id)
	if err == sql.ErrNoRows {
		// not matching contact found, time to create a new one ...
		fmt.Errorf("No contact found where both '%s' & '%s' exist.", phone, email)
		return -1, nil
	} else if err != nil {
		return -2, err
	}
	// contact id found, return int
	return contact_id, nil
}

func GetAddressID(db *sql.DB, name, street, city, state, zip string) (int, error) {
    var address_id int
    query := fmt.Sprintf("SELECT id FROM addresses WHERE name = ? AND street = ? AND city = ? AND state = ? AND zip = ?;")
    err := db.QueryRow(query, name, street, city, state, zip).Scan(&address_id)
    if err == sql.ErrNoRows {
        // no matching address
        fmt.Errorf("Adress does not exist in the system, create new ...")
        return -1, nil
    } else if err != nil {
        return -2, err
    }
    return address_id, nil
}

func DoesGroupExist(db *sql.DB, groupName, creatorUsername string) (int, error) {
    var status int
    if groupName == "" {
        return -1, nil
    } else if creatorUsername == "" {
        // FIXLOG 1
        return -2, nil
    } 
    groupQuery := "SELECT id FROM groups WHERE name = ?;"
    err := db.QueryRow(groupQuery, groupName).Scan(&status)
    if err == sql.ErrNoRows {
        // group name is not in db yet
        status = 1
    }
    return status, nil
}
