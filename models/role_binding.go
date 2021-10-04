package models

import (
	"gorm.io/gorm"
)

type RoleBinding struct {
	gorm.Model
	Name		string		`json:"name"`
	ApplicationID	uint
	Application	Application	`json:"application"`
	RoleID		uint
	Role		Role		`json:"role"`
	UserID		uint
	User		User		`json:"user"`
	GroupID		uint
	Group		Group		`json:"group"`
}

func GetRoleBindings(db *gorm.DB) ([]RoleBinding, error) {
        roleBindings := []RoleBinding{}
        query :=  db.Select("role_bindings.*").Group("role_bindings.id")
        if err := query.Find(&roleBindings).Error; err != nil {
                return roleBindings, err
        }

        return roleBindings, nil
}

func GetRoleBindingByID(db *gorm.DB, id string) (RoleBinding, bool, error) {
        roleBinding := RoleBinding{}
        query := db.Select("role_bindings.*").Group("role_bindings.id")
        err := query.Where("role_bindings.id = ?", id).First(&roleBinding).Error
        if err != nil {
                if err != gorm.ErrRecordNotFound {
                        return roleBinding, false, err
                } else {
                        return roleBinding, false, nil
                }
        }

        return roleBinding, true, nil
}

func CreateRoleBinding(db *gorm.DB, roleBinding *RoleBinding) error {
        if err := db.Create(&roleBinding).Error; err != nil {
                return err
        }

        return nil
}

func DeleteRoleBinding(db *gorm.DB, id string) error {
        var roleBinding RoleBinding
        if err := db.Where("id = ?", id).Delete(&roleBinding).Error; err != nil {
                return err
        }

        return nil
}

func UpdateRoleBinding(db *gorm.DB, roleBinding *RoleBinding) error {
        if err := db.Save(&roleBinding).Error; err != nil {
                return err
        }

        return nil
}

