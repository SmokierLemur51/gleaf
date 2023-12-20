/* 
    Things you can get from category.go

    Category Type 
        Methods
        - InsertCategory(db) -> Inserts a new category after calling CheckExistingCategory()
        - UpdateCategory(db) -> Updates category
        - RunReport(db, period) -> With period being day, week etc. Generates reports from that category
    Functions 
    - CheckExistingCategory(db, category) -> Returns true if exists, false if not
    - CheckCategoryFields(*Category) -> Check that required fields exist
*/
package data
import (
    "database/sql"
    "log"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)    

type ServiceCategory struct {
    Id int `db:"id"`
    Category string `db:"category"`
    Description string `db:"description"`
}

func CheckExistingCategory(db *sql.DB, category string) (bool, error) {
    // returns true if it exists
    var count int
    rows, err := db.Query("SELECT COUNT(*) FROM service_categories WHERE category = ?", category)
    if err != nil {
        return true, err
    }
    defer rows.Close()
    for rows.Next() {
        if err := rows.Scan(&count); err != nil {
            return false, err
        }
    }
    if count > 0 {
        return true, err
    }
    // if not
    return false, nil
}

func FindCategoryId(db *sql.DB, category string) string {
    rows, err := db.Query("SELECT id FROM service_categories WHERE category = ?", category)
    if err != nil {
        log.Println(err)
    }
    var c string
    for rows.Next() {
        err := rows.Scan(&c)
        if err != nil {
            log.Fatal(err)
            return ""
        }
    }
    return c
}

func (s *ServiceCategory) InsertCategory(db *sql.DB) {
    var execute bool
    var err error
    execute, err = CheckExistingService(db, s.Category)
    if err != nil {
        log.Println(err)
        return
    }
    // remember, the check existing returns true if the product already exists, so it skips
    switch execute {
    case false:
        _, err := db.Exec(
            "INSERT INTO service_categories (category, description) VALUES (?,?)",
            s.Category, s.Description,
        )
        if err != nil {
            log.Fatal(err)
        }
    case true:
        fmt.Printf("Category %s already exists.\n", s.Category)
    }
}