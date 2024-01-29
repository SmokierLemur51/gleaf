package models
/*
File: statcode.go
*/
import (
  "gorm.io/gorm"
)

func CreateStatCode(status, description string) StatusCode {
  return StatusCode{
    StatCode: status, StatDescription: description,
  }
}

func (sc *StatusCode) InsertStatusCode(db *gorm.DB) error {
  if result := db.Create(sc); result.Error != nil {
    return result.Error
  }
  return nil
}

func UpdateStatusCode(db *gorm.DB) error { return nil }

func DeleteStatusCode(db *gorm.DB) error { return nil }

func QueryStatCodeID(db *gorm.DB, status string) (uint, error) {
  var stat StatusCode 
  result := db.Where("stat_code = ?", status).First(&stat)
  if result.Error != nil {
    return 0, result.Error
  }
  return stat.ID, nil
}


