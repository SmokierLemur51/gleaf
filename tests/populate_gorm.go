package tests

import (
  "github.com/SmokierLemur51/gleaf/data"
  "gorm.io/gorm"
)

func CreateModels(db *gorm.DB) {
  db.AutoMigrate(
    &data.StatusCode{}, 
    &data.ServiceCategory{},
    &data.Service{}, 
    &data.ClientScore{},
    &data.GroupScore{}, 
    &data.Address{},
    &data.Client{},
    &data.Group{},
    &data.Estimate{},
    &data.GroupEstimate{},
    &data.Booking{},
    &data.GroupBooking{},
    &data.ChristmasCardMailingList{},
  )
}
  


func PopulateStatusCodes() []data.StatusCode {
  return []data.StatusCode{
    {StatusCode:"Active", Description:"Actively listed, promoted, or sold."},  
    {StatusCode:"Inactive", Description:"Not listed publicly."},
    {StatusCode:"Promotion", Description:"Promoted at a discounted price."},
    {StatusCode:"Sale", Description:"Discounted price."},
  }
} 

func PopulateServiceCategory() []data.ServiceCategory {
  return []data.ServiceCategory{}
}

func PopulateService() []data.Service {
  return []data.Service{}
}

func PopulateClientScore() []data.ClientScore {
  return []data.ClientScore{}
}


