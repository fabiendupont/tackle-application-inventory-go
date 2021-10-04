package models

import (
	"gorm.io/gorm"
)

type TagType struct {
	gorm.Model
	Name	string	`json:"name"`
	Colour	string	`json:"colour"`
	Rank	uint	`json:"rank"`
}

func GetTagTypes(db *gorm.DB) ([]TagType, error) {
	tagTypes := []TagType{}
	query :=  db.Select("tag_types.*").Group("tag_types.id")
	if err := query.Find(&tagTypes).Error; err != nil {
		return tagTypes, err
	}

	return tagTypes, nil
}

func GetTagTypeByID(db *gorm.DB, id string) (TagType, bool, error) {
	tagType := TagType{}
	query := db.Select("tag_types.*").Group("tag_types.id")
	err := query.Where("tag_types.id = ?", id).First(&tagType).Error
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

