package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name		string	`json:"name" gorm:"notnull" binding:"required"`
	TagTypeID	uint	`json:"tag_type_id" gorm:"notnull" binding:"required"`
	TagType		TagType
}

func GetTags(db *gorm.DB) ([]Tag, error) {
	tags := []Tag{}
	if err := db.Joins("TagType").Find(&tags).Error; err != nil {
		return tags, err
	}

	return tags, nil
}

func GetTagByID(db *gorm.DB, id string) (Tag, bool, error) {
	tag := Tag{}
	err := db.Joins("TagType").First(&tag, "tags.id = ?", id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return tag, false, err
		} else {
			return tag, false, nil
		}
	}

	return tag, true, nil
}

func CreateTag(db *gorm.DB, tag *Tag) error {
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTag(db *gorm.DB, id string) error {
	var tag Tag
	if err := db.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTag(db *gorm.DB, tag *Tag) error {
	if err := db.Save(&tag).Error; err != nil {
		return err
	}

	return nil
}

