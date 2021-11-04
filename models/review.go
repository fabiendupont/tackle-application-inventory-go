package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	BusinessCriticality	uint		`json:"business_criticality" gorm:"notnull" binding:"required"`
	EffortEstimate		string		`json:"effort_estimate" gorm:"notnull" binding:"required"`
	ProposedAction		string		`json:"proposed_action" gorm:"notnull" binding:"required"`
	WorkPriority		uint		`json:"work_priority" gorm:"notnull" binding:"required"`
	ApplicationID		uint		`json:"application_id"`
	Application		Application
}

func GetReviews(db *gorm.DB) ([]Review, error) {
	reviews := []Review{}
	query :=  db.Select("reviews.*").Group("reviews.id")
	if err := query.Find(&reviews).Error; err != nil {
		return reviews, err
	}

	return reviews, nil
}

func GetReviewByID(db *gorm.DB, id string) (Review, bool, error) {
	review := Review{}
	query := db.Select("reviews.*").Group("reviews.id")
	err := query.Where("reviews.id = ?", id).First(&review).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return review, false, err
		} else {
			return review, false, nil
		}
	}

	return review, true, nil
}

func CreateReview(db *gorm.DB, review *Review) error {
	if err := db.Create(&review).Error; err != nil {
		return err
	}

	return nil
}

func DeleteReview(db *gorm.DB, id string) error {
	var review Review
	if err := db.Where("id = ?", id).Delete(&review).Error; err != nil {
		return err
	}

	return nil
}

func UpdateReview(db *gorm.DB, review *Review) error {
	if err := db.Save(&review).Error; err != nil {
		return err
	}

	return nil
}

