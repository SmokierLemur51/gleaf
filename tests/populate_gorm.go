package tests

import (
  "github.com/SmokierLemur51/gleaf/models"
  "gorm.io/gorm"
)

func CreateModels(db *gorm.DB) {
  db.AutoMigrate(
    &models.StatusCode{}, 
    &models.ServiceCategory{},
    &models.Service{}, 
    &models.ClientScore{},
    &models.GroupScore{}, 
    &models.Address{},
    &models.Client{},
    &models.Group{},
    &models.Estimate{},
    &models.GroupEstimate{},    
    &models.Booking{},
    &models.GroupBooking{},
    &models.ChristmasCardMailingList{},
  )
}
  


func PopulateStatusCodes(db *gorm.DB) error {
    var statCodes = []*models.StatusCode{
        {StatCode: "active", StatDescription: "actively being sold"},
        {StatCode: "inactive", StatDescription: "currently inactive"},
        {StatCode: "sale", StatDescription: "temporarily discounting price"},
        {StatCode: "promotion", StatDescription: "trying to sell this, lowering price and actively pushing to customers"},
        {StatCode: "discontinued", StatDescription: "no longer open for bookings"},
        {StatCode: "seasonal", StatDescription: "seasonal promotion, ex: leaf removal in the fall"},
    }
    db.Create(statCodes)
    
    return nil 
}

func PopulateBookingStatusCodes(db *gorm.DB) error {
    var statCodes = []*models.StatusCode{
        {StatCode: "not contacted", StatDescription: "not yet contacted, does not include leaving a message"},
        {StatCode: "contacted", StatDescription: "spoke with potiential customer, not scheduled or rejected"},
        {StatCode: "accepted", StatDescription: "spoke with customer, offer accepted and scheduled"},
        {StatCode: "rejected", StatDescription: "spoke with customer, was not interested in offer"},
    }
    result := db.Create(statCodes)
    if result.Error != nil {
        return result.Error
    }    
    return nil
}


func PopulateServiceCategories(db *gorm.DB) error {
  active, err := models.QueryStatCodeID(db, "active")
  if err != nil {
    panic(err)
  }
  var servCats = []*models.ServiceCategory{
    {StatusCodeID: active, Category: "Moving", 
      AdminInformation: "Moving related services.", 
      PublicInformation: "Services related to moving, in or out, we are here to help you."},
    {StatusCodeID: active, Category: "Residential", 
      AdminInformation: "Customer house cleaning, no commercial offers. Wide umbrella to define general jobs.", 
      PublicInformation: 
      "Whether its help catching up after a busy week, toddlers gone wild, a grocery pickup, or a party you don't have time to prepare for, we've got your back!",
    },
    {StatusCodeID: active, Category: "Residential Exterior",
      AdminInformation: "Non commercial exterior cleaning.",
      PublicInformation: "Windows, siding, decks, gutters, leaf removal, driveways and garages. Almost anything except for junk removal.",
    },
    {StatusCodeID: active, Category: "Eco-Friendly", 
      AdminInformation: "Green cleaning solutions, its cheaper and healthier",
      PublicInformation: "Greenleaf Cleanings specialty! Ditch those harmful chemicals to protect your health and the environment!",
    },
    {StatusCodeID: active, Category: "Group Cleaning",
      AdminInformation: "Our attempt to stand out from the crowd. Share booking links to friends if you have an account and recieve discounts based on the number of houses booked.",
      PublicInformation: "Create a group and book with your friends, family or neighbors! We offer discounts based off of the size, distance, and history of booking with us.",
    },
    {StatusCodeID: active, Category: "Commercial",
      AdminInformation: "Office, warehouse, or moving company bookings. Etc...",
      PublicInformation: "Commercial cleaning solutions, contact us to get a free estimate.",
    },
  }
  result := db.Create(servCats)
  if result.Error != nil {
    return result.Error
  }  
  return nil
}

func PopulateServiceCategory() []models.ServiceCategory {
  return []models.ServiceCategory{}
}

func PopulateService() []models.Service {
  return []models.Service{}
}

func PopulateClientScore() []models.ClientScore {
  return []models.ClientScore{}
}
