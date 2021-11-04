package models

import (
	"gorm.io/gorm"
)

type SourceRepository struct {
	gorm.Model
	Type		string		`json:"name" gorm:"notnull" binding:"required" validate:"oneof= git svn"`
	URL		string		`json:"url" gorm:"notnull" binding:"required"`
	Branch		string		`json:"branch" gorm:"notnull" binding:"required"`
	ApplicationID	uint		`json:"application_id"`
	Application	Application
}

func GetSourceRepositories(db *gorm.DB) ([]SourceRepository, error) {
	source_repositorys := []SourceRepository{}
	query :=  db.Select("source_repositorys.*").Group("source_repositorys.id")
	if err := query.Find(&source_repositorys).Error; err != nil {
		return source_repositorys, err
	}

	return source_repositorys, nil
}

func GetSourceRepositoryByID(db *gorm.DB, id string) (SourceRepository, bool, error) {
	source_repository := SourceRepository{}
	query := db.Select("source_repositorys.*").Group("source_repositorys.id")
	err := query.Where("source_repositorys.id = ?", id).First(&source_repository).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return source_repository, false, err
		} else {
			return source_repository, false, nil
		}
	}

	return source_repository, true, nil
}

func CreateSourceRepository(db *gorm.DB, source_repository *SourceRepository) error {
	if err := db.Create(&source_repository).Error; err != nil {
		return err
	}

	return nil
}

func DeleteSourceRepository(db *gorm.DB, id string) error {
	var source_repository SourceRepository
	if err := db.Where("id = ?", id).Delete(&source_repository).Error; err != nil {
		return err
	}

	return nil
}

func UpdateSourceRepository(db *gorm.DB, source_repository *SourceRepository) error {
	if err := db.Save(&source_repository).Error; err != nil {
		return err
	}

	return nil
}

