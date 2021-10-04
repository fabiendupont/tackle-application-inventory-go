package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name	string	`json:"name"`
}

func GetRoles(db *gorm.DB) ([]Role, error) {
        roles := []Role{}
        query :=  db.Select("roles.*").Group("roles.id")
        if err := query.Find(&roles).Error; err != nil {
                return roles, err
        }

        return roles, nil
}

func GetRoleByID(db *gorm.DB, id string) (Role, bool, error) {
        role := Role{}
        query := db.Select("roles.*").Group("roles.id")
        err := query.Where("roles.id = ?", id).First(&role).Error
        if err != nil {
                if err != gorm.ErrRecordNotFound {
                        return role, false, err
                } else {
                        return role, false, nil
                }
        }

        return role, true, nil
}

func CreateRole(db *gorm.DB, role *Role) error {
        if err := db.Create(&role).Error; err != nil {
                return err
        }

        return nil
}

func DeleteRole(db *gorm.DB, id string) error {
        var role Role
        if err := db.Where("id = ?", id).Delete(&role).Error; err != nil {
                return err
        }

        return nil
}

func UpdateRole(db *gorm.DB, role *Role) error {
        if err := db.Save(&role).Error; err != nil {
                return err
        }

        return nil
}

