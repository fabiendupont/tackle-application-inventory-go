package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string	`json:"username" gorm:"notnull,unique,index" binding:"required,alphanum,min=6,max=32"`
	DisplayName	string	`json:"displayName" gorm:"notnull" binding:"required"`
	Email		string	`json:"email" gorm:"notnull" binding:"required,email"`
	Groups		[]Group	`json:"groups" gorm:"many2many:user_groups"`
}

func GetUsers(db *gorm.DB) ([]User, error) {
        users := []User{}
        query :=  db.Select("users.*").Group("users.id")
        if err := query.Find(&users).Error; err != nil {
                return users, err
        }

        return users, nil
}

func GetUserByID(db *gorm.DB, id string) (User, bool, error) {
        user := User{}
        query := db.Select("users.*").Group("users.id")
        err := query.Where("users.id = ?", id).First(&user).Error
        if err != nil {
                if err != gorm.ErrRecordNotFound {
                        return user, false, err
                } else {
                        return user, false, nil
                }
        }

        return user, true, nil
}

func CreateUser(db *gorm.DB, user *User) error {
        if err := db.Create(&user).Error; err != nil {
                return err
        }

        return nil
}

func DeleteUser(db *gorm.DB, id string) error {
        var user User
        if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
                return err
        }

        return nil
}

func UpdateUser(db *gorm.DB, user *User) error {
        if err := db.Save(&user).Error; err != nil {
                return err
        }

        return nil
}

