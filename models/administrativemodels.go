package models

type AdministrativePageData struct {
	Page 				string
	Title   			string
	ServiceCategories   []ServiceCategory
	Services 			[]Service
	IncompleteContactRequests []ContactRequest
	IncompleteBookings  []Bookings
	FinancialData		Finances
}


type Finances struct {
	GrossIncome		float32
	Taxes  			float32
	MonthlyExpenses		[]MonthlyExpenses
}

type MonthlyExpenses struct {
	Advertising		float32
	Wages			float32
	Gas  			float32
	Materials   	float32
	OfficeSupplies  float32
	Hosting			float32
	Meals 			float32
	Rent 			float32
}

type Bookings struct {}

