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
    "log"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

// status 1 active, status 2 inactive
type Service struct {
    Id int             `db:"id"`
    CategoryId int     `db:"category_id"`
    Category ServiceCategory
    Status int         `db:"status_id"`
    Service string     `db:"service"`
    Description string `db:"description"`
    Selling float64    `db:"selling"`
}

func CheckServiceFields(s *Service) (*Service, error) {return s, nil}

func CheckExistingService(db *sql.DB, service string) (bool, error) {
    // returns true if it exists
    var count int
    rows, err := db.Query("SELECT COUNT(*) FROM services WHERE service = ? ", service)
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

func (s Service) InsertService(db *sql.DB) {
    var execute bool
    var err error
    execute, err = CheckExistence(db, "services", "service", s.Service)
    if err != nil {
        log.Println(err)
        return
    }
    // remember, the check existing returns true if the product already exists, so it skips
    switch execute {
    case false:
        _, err := db.Exec(
            "INSERT INTO services (category_id, status_id, service, description, selling) VALUES (?,?,?,?,?)",
            s.CategoryId, s.Status, s.Service, s.Description, s.Selling,           
        )
        if err != nil {
            log.Fatal(err)
        }
    case true:
        fmt.Printf("Service %s already exists.\n", s.Service)
    }
}

func (s Service) UpdateSellingPrice(db *sql.DB, selling float64) {}

func (s Service) RunReport(db *sql.DB, period string) {
    // period will be day, week, month, year, all
    switch period {
        case "day":
        fmt.Println(period)
        
        case "week":
        fmt.Println(period)
        
        case "month":
        fmt.Println(period)
        
        case "quarter":
        fmt.Println(period)
        
        case "year":
        fmt.Println(period)
        
        case "all":
        fmt.Println(period)
    }
}

// General functions
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

func LoadServicesByCategory(db *sql.DB, category string) ([]Service, error) {
    var services []Service
    rows, err := db.Query(
        "SELECT id, category_id, status_id, service, description, selling FROM services WHERE category_id = ?",
        FindCategoryId(db, category),
    )
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        var s Service
        err := rows.Scan(&s.Id, &s.CategoryId, &s.Status, &s.Service, &s.Description, &s.Selling)
        if err != nil {
            log.Fatal(err)
        }
        services = append(services, s)
        log.Printf("Added service: %s to %s slice.", s.Service, category)
    }
    return services, nil 
}

