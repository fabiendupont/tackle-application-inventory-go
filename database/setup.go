package database

import (
        "github.com/fabiendupont/tackle-hub/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
)

var jobFunctions = [...]models.JobFunction {
	{gorm.Model{}, "Business Analyst"},
	{gorm.Model{}, "Business Service Owner / Manager"},
	{gorm.Model{}, "Consultant"},
	{gorm.Model{}, "DBA"},
	{gorm.Model{}, "Developer / Software Engineer"},
	{gorm.Model{}, "IT Operations"},
	{gorm.Model{}, "Program Manager"},
	{gorm.Model{}, "Project Manager"},
	{gorm.Model{}, "Service Owner"},
	{gorm.Model{}, "Solution Architect"},
	{gorm.Model{}, "System Administrator"},
	{gorm.Model{}, "Test Analyst / Manager"},
}

var tagTypes = [...]models.TagType {
//	{
//		gorm.Model{},
//		"Application Type",
//		6,
//		"#ec7a08",
//		[...]models.Tag {
//			{gorm.Model{}, "COTS"},
//			{gorm.Model{}, "In house"},
//			{gorm.Model{}, "SaaS"},
//		},
//	},
//	{gorm.Model{}, "Data Center", 5, "#2b9af3"},	// id = 2
//	{gorm.Model{}, "Database", 4, "#6ec664"},		// id = 3
//	{gorm.Model{}, "Language", 1, "#009596"},		// id = 4
//	{gorm.Model{}, "Operating System", 2, "#a18fff"},	// id = 5
//	{gorm.Model{}, "Runtime", 3, "#7d1007"},		// id = 6
}

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(sqlite.Open("tackle.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Application{},
		&models.BinaryRepository{},
		&models.BusinessService{},
		&models.Group{},
		&models.JobFunction{},
		&models.Review{},
		&models.Role{},
		&models.RoleBinding{},
		&models.SourceRepository{},
		&models.Tag{},
		&models.TagType{},
		&models.User{},
	)

	LoadJobFunctions(db)
//	LoadTagTypes(db)
//	LoadTags(db)

	DB = db
}

// Return gorm.DB pointer
func GetBD() *gorm.DB {
	return DB
}

// Load job functions in the database
func LoadJobFunctions(db *gorm.DB) {
	for _, jobFunction := range(jobFunctions) {
		models.CreateJobFunction(db, &jobFunction)
	}
}

// Preload tag types in the database
func LoadTagTypes(db *gorm.DB) {
	for _, tagType := range(tagTypes) {
		models.CreateTagType(db, &tagType)
	}
}

// Preload tags in the database
//func LoadTags(db *gorm.DB) {
// TODO: Do not hard code tag types
//	tagType := models.GetTagTypeByName("Application Type")

//var tags = [...]models.Tag {
//	// TagType: Application Type
//	models.Tag{gorm.Model{}, "COTS", 1},
//	models.Tag{gorm.Model{}, "In hose", 1},
//	models.Tag{gorm.Model{}, "SaaS", 1},
//	// TagType: Data Center
//	models.Tag{gorm.Model{}, "Boston (USA)", 2},
//	models.Tag{gorm.Model{}, "London (UK)", 2},
//	models.Tag{gorm.Model{}, "Paris (FR)", 2},
//	models.Tag{gorm.Model{}, "Sydney (AU)", 2},
//	// TagType: Database
//	models.Tag{gorm.Model{}, "DB2", 3},
//	models.Tag{gorm.Model{}, "MongoDB", 3},
//	models.Tag{gorm.Model{}, "Oracle", 3},
//	models.Tag{gorm.Model{}, "PostgreSQL", 3},
//	models.Tag{gorm.Model{}, "SQL Server", 3},
//	// TagType: Language
//	models.Tag{gorm.Model{}, "C# ASP .Net", 4},
//	models.Tag{gorm.Model{}, "C++", 4},
//	models.Tag{gorm.Model{}, "COBOL", 4},
//	models.Tag{gorm.Model{}, "Java", 4},
//	models.Tag{gorm.Model{}, "Javascript", 4},
//	models.Tag{gorm.Model{}, "Python", 4},
//	// TagType: Operating System
//	// TODO: Use osinfo identifiers for better link with Kubevirt
//	models.Tag{gorm.Model{}, "RHEL 8", 5},
//	models.Tag{gorm.Model{}, "Windows Server 2016", 5},
//	models.Tag{gorm.Model{}, "Z/OS", 5},
//	// TagType: Runtime
//	models.Tag{gorm.Model{}, "EAP", 6},
//	models.Tag{gorm.Model{}, "JWS", 6},
//	models.Tag{gorm.Model{}, "Quarkus", 6},
//	models.Tag{gorm.Model{}, "Spring Boot", 6},
//	models.Tag{gorm.Model{}, "Tomcat", 6},
//	models.Tag{gorm.Model{}, "WebLogic", 6},
//	models.Tag{gorm.Model{}, "WebSphere", 6},
//}

//	for _, tag := range(tags) {
//		models.CreateTag(db, &tag)
//	}
//}
