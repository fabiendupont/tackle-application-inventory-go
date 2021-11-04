package models

import (
	"gorm.io/gorm"
)


type Application struct {
	gorm.Model
	Name			string		`json:"name" gorm:"notnull" binding:"required"`
	Description		string		`json:"description"`
	Comments		string		`json:"comments"`
	BusinessServiceID	uint		`json:"business_service_id" gorm:"notnull" binding:"required"`
	BusinessService		BusinessService
	DependsOn		[]*Application	`json:"depends_on" gorm:"many2many:application_dependencies"`
	Tags			[]Tag		`json:"tags" gorm:"many2many:application_tags"`
//	ReviewID		uint		`json:"review_id"`
//	Review			Review		`json:"review"`
}

func GetApplications(db *gorm.DB) ([]Application, error) {
	applications := []Application{}
	query :=  db.Select("applications.*").Group("applications.id")
	if err := query.Find(&applications).Error; err != nil {
		return applications, err
	}

	return applications, nil
}

func GetApplicationByID(db *gorm.DB, id string) (Application, bool, error) {
	application := Application{}
	query := db.Select("applications.*").Group("applications.id")
	err := query.Where("applications.id = ?", id).First(&application).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return application, false, err
		} else {
			return application, false, nil
		}
	}

	return application, true, nil
}

func CreateApplication(db *gorm.DB, application *Application) error {
	if err := db.Create(&application).Error; err != nil {
		return err
	}

	return nil
}

func DeleteApplication(db *gorm.DB, id string) error {
	var application Application
	if err := db.Where("id = ?", id).Delete(&application).Error; err != nil {
		return err
	}

	return nil
}

func UpdateApplication(db *gorm.DB, application *Application) error {
	if err := db.Save(&application).Error; err != nil {
		return err
	}

	return nil
}
