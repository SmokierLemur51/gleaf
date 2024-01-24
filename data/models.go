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
  Status_Code StatusCode 
  Category string 
  AdminInformation string
  PublicInformation string
}

type Service struct {
  gorm.Model
  Status_Code StatusCode
  Category ServiceCategory
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
  Client_Score ClientScore
  Address Address // one to many? 
  Name string
  Email *string
  Phone string
}

type Group struct {
  gorm.Model
  SecretIdentity string // random string generated for endpoint string
  URL string
  Name string
  Creator Client
  Clients []Client
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
