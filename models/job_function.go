package models

import (
	"gorm.io/gorm"
)

type JobFunction struct {
	gorm.Model
	Name	string	`json:"name" gorm:"notnull,unique"`
}

func GetJobFunctions(db *gorm.DB) ([]JobFunction, error) {
        job_functions := []JobFunction{}
        query :=  db.Select("job_functions.*").Group("job_functions.id")
        if err := query.Find(&job_functions).Error; err != nil {
                return job_functions, err
        }

        return job_functions, nil
}

func GetJobFunctionByID(db *gorm.DB, id string) (JobFunction, bool, error) {
        job_function := JobFunction{}
        query := db.Select("job_functions.*").Group("job_functions.id")
        err := query.Where("job_functions.id = ?", id).First(&job_function).Error
        if err != nil {
                if err != gorm.ErrRecordNotFound {
                        return job_function, false, err
                } else {
                        return job_function, false, nil
                }
        }

        return job_function, true, nil
}

func CreateJobFunction(db *gorm.DB, job_function *JobFunction) error {
        if err := db.Create(&job_function).Error; err != nil {
                return err
        }

        return nil
}

func DeleteJobFunction(db *gorm.DB, id string) error {
        var job_function JobFunction
        if err := db.Where("id = ?", id).Delete(&job_function).Error; err != nil {
                return err
        }

        return nil
}

func UpdateJobFunction(db *gorm.DB, job_function *JobFunction) error {
        if err := db.Save(&job_function).Error; err != nil {
                return err
        }

        return nil
}

