/*
File: models.go
Creator: Logan Lee @SmokierLemur51

Gorm models defined here.
*/
package data

import (
  "gorm.io/gorm"
)

type StatusCode struct {
  gorm.Model
  StatCode string 
  StatDescription string
}

type ServiceCategory struct {
  gorm.Model
  StatusCodeID uint `gorm:"not null"`
  Status_Code StatusCode `gorm:"foreignKey:StatusCodeID"`
  Category string 
  AdminInformation string
  PublicInformation string
  Services []Service `gorm:"foreignKey:ServiceCategoryID"`
}

type Service struct {
  gorm.Model
  StatusCodeID uint `gorm:"not null"`
  Status_Code StatusCode `gorm:"foreignKey:StatusCodeID"`
  ServiceCategoryID uint `gorm:"not null"`
  Category ServiceCategory `gorm:"foreignKey:ServiceCategoryID"`
  Service string
  AdminDescription string
  PublicDescription string
  Cost float64
  SellingPrice float64
}

type ClientScore struct {
  gorm.Model
  RequiredBookings uint 
  ScoringLevel string
}

type GroupScore struct {
  gorm.Model
  RequiredBookings uint
  ScoringLevel string 
}

type Address struct {
  gorm.Model
  Street1 string
  Street2 *string
  City string
  State string
  Zip string
}

type Client struct {
  gorm.Model
  ClientScoreID uint `gorm:"not null"`
  Client_Score ClientScore `gorm:"foreignKey:ClientScoreID"`
  AddressID uint `gorm:"not null"`
  Address Address `gorm:"foreignKey:AddressID"` 
  GroupID uint `gorm:"nullable"`
  Group GroupInterface
  Name string
  Email *string
  Phone string
}

func (c *Client) IsGroup() bool {
  if c.GroupID {
    return true
  }
  return false
}

type GroupInterface interface {
  IsGroup()
}

type Group struct {
  gorm.Model
  SecretIdentity string // random string generated for endpoint string
  URL string
  Name string
  CreatorID uint `gorm:"not null"`
  Creator Client `gorm:"foreignKey:CreatorID"`
  Clients []Client `gorm:"foreignKey:GroupID"`
}

type Estimate struct {
  gorm.Model
}

type GroupEstimate struct {
  gorm.Model
}

type Booking struct {
  gorm.Model
}


type GroupBooking struct {
  // Should this be multiple bookings with a group foreign key
  // or should it be a different entity all together?
  gorm.Model
}

type ChristmasCardMailingList struct {
  gorm.Model
}

func CreateModels(db *gorm.DB) {
  db.AutoMigrate(
    &StatusCode{}, &ServiceCategory{}, &Service{}, &ClientScore{}, &GroupScore{},
    &Address{}, &Client{}, &Group{}, &Estimate{}, &GroupEstimate{}, 
    &Booking{}, &GroupBooking{}, &ChristmasCardMailingList{},
  )
}
