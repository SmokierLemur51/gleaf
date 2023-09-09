package models 

type PageData struct {
	Page 		string
	Title		string
	UserHash 	[]byte
	Message 	string
	Services    []Service
}

type Service struct {
	ID			int16
	Type_ID		int16
	Type 		string
	Description string
	Cost 		float32
}

type User struct {
	ID int16
	Username string
	Password string

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




