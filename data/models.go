package data

import (
  "gorm.io/gorm"
)

type ServiceCode struct {
  gorm.Model
  _Status string 
  _Description string
}

type ServiceCategory struct {
  gorm.Model
}

type Service struct {
  gorm.Model
}
