package models

import (
	"gorm.io/gorm"
)

type TagType struct {
	gorm.Model
	Name	string	`json:"name" gorm:"notnull" binding:"required"`
	Rank	uint	`json:"rank"`
	Colour	string	`json:"colour"`
	Tags	[]Tag	`json:"tags"`
}

func GetTagTypes(db *gorm.DB) ([]TagType, error) {
	tagTypes := []TagType{}
	if err := db.Preload("Tags").Find(&tagTypes).Error; err != nil {
		return tagTypes, err
	}

	return tagTypes, nil
}

func GetTagTypeByID(db *gorm.DB, id string) (TagType, bool, error) {
	tagType := TagType{}
	err := db.Preload("Tags").First(&tagType, "tag_types.id = ?", id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return tagType, false, err
		} else {
			return tagType, false, nil
		}
	}

	return tagType, true, nil
}

func CreateTagType(db *gorm.DB, tagType *TagType) error {
	if err := db.Create(&tagType).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTagType(db *gorm.DB, id string) error {
	var tagType TagType
	if err := db.Where("id = ?", id).Delete(&tagType).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTagType(db *gorm.DB, tagType *TagType) error {
	if err := db.Save(&tagType).Error; err != nil {
		return err
	}

	return nil
}

