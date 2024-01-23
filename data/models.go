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

type ServicerviceCategory struct {
  gorm.Model
  StatusCode StatusCode 
  Category string 
  AdminInformation string
  PublicInformation string
}

type Service struct {
  gorm.Model
  StatusCode StatusCode
  Category Category
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
  ClientScore ClientScore
  Address Address // one to many? 
  Name string
  Email *string
  Phone string
}

type Group struct {
  gorm.Model
  SecretIdentity string // random string generated for endpoint string
  Name string
  Creator Client
  Clients []Client
}

type Estimate struct {}

type GroupEstimate struct {}

type Booking struct {}

type GroupBooking struct {} 
