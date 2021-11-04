package models

import (
	"gorm.io/gorm"
)

type BinaryRepository struct {
	gorm.Model
	Type		string		`json:"name" gorm:"notnull" binding:"required" validate:"oneof=mvn"`
	URL		string		`json:"url" gorm:"notnull" binding:"required"`
	Group		string		`json:"group" gorm:"notnull" binding:"required"`
	Artifact	string		`json:"artifact" gorm:"notnull" binding:"required"`
	Version		string		`json:"version" gorm:"notnull" binding:"required`
	ApplicationID	uint		`json:"application_id"`
	Application	Application
}

func GetBinaryRepositories(db *gorm.DB) ([]BinaryRepository, error) {
	binary_repositorys := []BinaryRepository{}
	query :=  db.Select("binary_repositorys.*").Group("binary_repositorys.id")
	if err := query.Find(&binary_repositorys).Error; err != nil {
		return binary_repositorys, err
	}

	return binary_repositorys, nil
}

func GetBinaryRepositoryByID(db *gorm.DB, id string) (BinaryRepository, bool, error) {
	binary_repository := BinaryRepository{}
	query := db.Select("binary_repositorys.*").Group("binary_repositorys.id")
	err := query.Where("binary_repositorys.id = ?", id).First(&binary_repository).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return binary_repository, false, err
		} else {
			return binary_repository, false, nil
		}
	}

	return binary_repository, true, nil
}

func CreateBinaryRepository(db *gorm.DB, binary_repository *BinaryRepository) error {
	if err := db.Create(&binary_repository).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBinaryRepository(db *gorm.DB, id string) error {
	var binary_repository BinaryRepository
	if err := db.Where("id = ?", id).Delete(&binary_repository).Error; err != nil {
		return err
	}

	return nil
}

func UpdateBinaryRepository(db *gorm.DB, binary_repository *BinaryRepository) error {
	if err := db.Save(&binary_repository).Error; err != nil {
		return err
	}

	return nil
}

