package models
/*
import (
  "fmt"
  "gorm.io/gorm"
)

func QueryStatCodeID(db *gorm.DB, status string) (uint, error) {
  var stat StatusCode 
  result := db.Where("stat_code = ?", status).First(&stat)
  if result.Error != nil {
    return 0, result.Error
  }
  return stat.ID, nil
}




func QueryAllCategoriesByStatus(db *gorm.DB, status string) ([]ServiceCategory, error) {
  return []ServiceCategory{}, nil
}

func QueryCategoryID(db *gorm.DB, category string) (uint, error) {
  var id uint
  if result := db.Where("category = ?", category).First(&id); result.Error != nil {
    return 0, result.Error
  }
  return id, nil
}

func (*cat ServiceCategory) CheckExistence(db *gorm.DB) bool {
  var cats []ServiceCategory 
  // load all categories 
  loaded := db.Find(&cats)
  // compare cat.Category against loaded
  for _, i := range cats {
    fmt.Printf("Category: %s\n", i.Category)
    if cat.Category == i.Category {
      fmt.Println("Shit")
    }
  }
  return false
}

*/
