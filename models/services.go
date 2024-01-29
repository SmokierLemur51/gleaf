package models
/*
import (
  "gorm.io/gorm"
)

func CreateService(statID, catID uint, service, adminInfo, publicInfo string, cost, selling float64) Service {
  return Service{
    StatusCodeID: statID, ServiceCategoryID: catID,
    Service: service, AdminInformation: adminInfo,
    PublicInformation: publicInfo, Cost: cost, Selling: selling,
  }
}

func (s *Service) InsertService(db *gorm.DB) error { 
  if result := db.Create(s); result.Error != nil {
    return result.Error
  }
  return nil 
}

func (s *Service) UpdateService() error { return nil }

func (s *Service) DeleteService() error { return nil }
*/
