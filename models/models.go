package models 


type PageData struct {
	Page 	string
	Title	string
	Message string
}

type ContactRequest struct {
	ID int16
	ContactInfo []ContactInfo
	Contacted   bool
}

type ContactInfo struct {
	Name	  string
	Email	  string
	Number    string
}

type Address struct {
	ID		  int16
	Street1   string
	Street2   string
	City 	  string
	State     string
	Zip		  string
}




