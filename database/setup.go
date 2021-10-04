package database

import (
        "github.com/fabiendupont/tackle-application-inventory-go/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(sqlite.Open("tackle.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Application{},
		&models.BusinessService{},
		&models.Group{},
		&models.Role{},
		&models.RoleBinding{},
		&models.Tag{},
		&models.TagType{},
		&models.User{},
	)
	DB = db
}

func GetBD() *gorm.DB {
	return DB
}
