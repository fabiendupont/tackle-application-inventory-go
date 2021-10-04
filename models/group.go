package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name	string	`json:"name"`
	Users	[]User	`json:"users" gorm:"many2many:user_groups"`
}

func GetGroups(db *gorm.DB) ([]Group, error) {
        groups := []Group{}
        query :=  db.Select("groups.*").Group("groups.id")
        if err := query.Find(&groups).Error; err != nil {
                return groups, err
        }

        return groups, nil
}

func GetGroupByID(db *gorm.DB, id string) (Group, bool, error) {
        group := Group{}
        query := db.Select("groups.*").Group("groups.id")
        err := query.Where("groups.id = ?", id).First(&group).Error
        if err != nil {
                if err != gorm.ErrRecordNotFound {
                        return group, false, err
                } else {
                        return group, false, nil
                }
        }

        return group, true, nil
}

func CreateGroup(db *gorm.DB, group *Group) error {
        if err := db.Create(&group).Error; err != nil {
                return err
        }

        return nil
}

func DeleteGroup(db *gorm.DB, id string) error {
        var group Group
        if err := db.Where("id = ?", id).Delete(&group).Error; err != nil {
                return err
        }

        return nil
}

func UpdateGroup(db *gorm.DB, group *Group) error {
        if err := db.Save(&group).Error; err != nil {
                return err
        }

        return nil
}

