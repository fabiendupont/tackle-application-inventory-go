package models

import (
	"gorm.io/gorm"
)

type JobFunctionBinding struct {
	gorm.Model
	Name			string		`json:"name"`
	BusinessServiceID	uint
	BusinessService		BusinessService	`json:"business_service"`
	JobFunctionID		uint
	JobFunction		JobFunction	`json:"job_function"`
	UserID			uint
	User			User		`json:"user"`
	GroupID			uint
	Group			Group		`json:"group"`
}

func GetJobFunctionBindings(db *gorm.DB) ([]JobFunctionBinding, error) {
        jobFunctionBindings := []JobFunctionBinding{}
	if err := db.Preload("BusinessService").Find(&jobFunctionBindings).Error; err != nil {
                return jobFunctionBindings, err
        }

        return jobFunctionBindings, nil
}

func GetJobFunctionBindingByID(db *gorm.DB, id string) (JobFunctionBinding, bool, error) {
        jobFunctionBinding := JobFunctionBinding{}
        err := db.Preload("BusinessService").First(&jobFunctionBinding, "job_function_bindings.id = ?", id).Error
        if err != nil {
                if err != gorm.ErrRecordNotFound {
                        return jobFunctionBinding, false, err
                } else {
                        return jobFunctionBinding, false, nil
                }
        }

        return jobFunctionBinding, true, nil
}

func CreateJobFunctionBinding(db *gorm.DB, jobFunctionBinding *JobFunctionBinding) error {
        if err := db.Create(&jobFunctionBinding).Error; err != nil {
                return err
        }

        return nil
}

func DeleteJobFunctionBinding(db *gorm.DB, id string) error {
        var jobFunctionBinding JobFunctionBinding
        if err := db.Where("id = ?", id).Delete(&jobFunctionBinding).Error; err != nil {
                return err
        }

        return nil
}

func UpdateJobFunctionBinding(db *gorm.DB, jobFunctionBinding *JobFunctionBinding) error {
        if err := db.Save(&jobFunctionBinding).Error; err != nil {
                return err
        }

        return nil
}

