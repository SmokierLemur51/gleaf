package data

import (
	"time"
)

// * * * * * * * * * * * * * * * Service Related  * * * * * * * * * * * * * * * *

type ServiceCategory struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type Service struct {
	ID           int `db:"id"`
	Type_ID      int `db:"category_id"`
	CategoryName string
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Cost         float32 `db:"cost"`
	Status       bool    `db:"status"`
	ImageURL     string  `db:"image_url"`
}

// * * * * * * * * * * * * * * * Contact Related  * * * * * * * * * * * * * * * *

type User struct {
	ID       int16
	Username string
	Password string
}

type ContactRequest struct {
	ID          int16
	ContactInfo []ContactInfo
	Contacted   bool
}

type ContactInfo struct {
	Name   string
	Email  string
	Number string
}

type Address struct {
	ID      int16
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
}

type CurrentUser struct {
	ID      int
	GroupID int
}

type PaymentInformation struct {
	PaymentID     int
	Total         float32
	GroupDiscount float32
	Discount      float32
	Paid          bool
	PaymentDate   time.Time
}

type Finances struct {
	GrossIncome     float32
	Taxes           float32
	MonthlyExpenses []MonthlyExpenses
}

type MonthlyExpenses struct {
	Advertising    float32
	Wages          float32
	Gas            float32
	Materials      float32
	OfficeSupplies float32
	Hosting        float32
	Meals          float32
	Rent           float32
}

type Bookings struct{}
