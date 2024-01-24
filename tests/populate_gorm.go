package tests

import (
  "github.com/SmokierLemur51/gleaf/data"
  "gorm.io/gorm"
)

func CreateModels(db *gorm.DB) {
  db.AutoMigrate(&data.StatusCode{})
  db.AutoMigrate(&data.ServiceCategory{})
  db.AutoMigrate(&data.Service{})
  db.AutoMigrate(&data.ClientScore{})
  db.AutoMigrate(&data.GroupScore{})
  db.AutoMigrate(&data.Address{})
  db.AutoMigrate(&data.Client{})
  db.AutoMigrate(&data.Group{})
  db.AutoMigrate(&data.Estimate{})
  db.AutoMigrate(&data.GroupEstimate{})
  db.AutoMigrate(&data.Booking{})
  db.AutoMigrate(&data.GroupBooking{})
  db.AutoMigrate(&data.ChristmasCardMailingList{})
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


