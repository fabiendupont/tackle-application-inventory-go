package models

import (
	"gorm.io/gorm"
)

type BusinessService struct {
	gorm.Model
	Name		string	`json:"name"`
	Description	string	`json:Description"`
	UserID		uint
	User		User	`json:"owner"`
}

func GetBusinessServices(db *gorm.DB) ([]BusinessService, error) {
	businessServices := []BusinessService{}
	query :=  db.Select("business_services.*").Group("business_services.id")
	if err := query.Find(&businessServices).Error; err != nil {
		return businessServices, err
	}

	return businessServices, nil
}

func GetBusinessServiceByID(db *gorm.DB, id string) (BusinessService, bool, error) {
	businessService := BusinessService{}
	query := db.Select("business_services.*").Group("business_services.id")
	err := query.Where("business_services.id = ?", id).First(&businessService).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return businessService, false, err
		} else {
			return businessService, false, nil
		}
	}

	return businessService, true, nil
}

func CreateBusinessService(db *gorm.DB, businessService *BusinessService) error {
	if err := db.Create(&businessService).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBusinessService(db *gorm.DB, id string) error {
	var businessService BusinessService
	if err := db.Where("id = ?", id).Delete(&businessService).Error; err != nil {
		return err
	}

	return nil
}

func UpdateBusinessService(db *gorm.DB, businessService *BusinessService) error {
	if err := db.Save(&businessService).Error; err != nil {
		return err
	}

	return nil
}

