package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetReviews godoc
// @summary Get all reviews.
// @description get all reviews.
// @tags get_reviews
// @produce json
// @success 200 {object} models.Review
// @router /application-inventory/review [get]
func GetReviews(c *gin.Context) {
	reviews, err := models.GetReviews(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReview godoc
// @summary Get a single review by its id.
// @description get a single review by its id.
// @tags get_review
// @produce json
// @success 200 {object} models.Review
// @router /application-inventory/review/:id [get]
// @param id path integer true "Review id"
func GetReview(c *gin.Context) {
	id := c.Params.ByName("id")
	review, exists, err := models.GetReviewByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no review in db")
		return
	}

	c.JSON(http.StatusOK, review)
}

// CreateReview godoc
// @summary Create a single review.
// @description Create a single review.
// @tags create_review
// @accept json
// @produce json
// @success 200 {object} models.Review
// @router /application-inventory/review [post]
// @param review body models.Review true "Review data"
func CreateReview(c *gin.Context) {
	review := models.Review{}
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateReview(database.DB, &review); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, review)
}

// DeleteReview godoc
// @summary Delete a single review by its id.
// @description delete a single review by its id.
// @tags delete_review
// @success 200 {object} models.Review
// @router /application-inventory/review/:id [delete]
// @param id path integer true "Review id"
func DeleteReview(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetReviewByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no review in db")
		return
	}

	if err = models.DeleteReview(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateReview godoc
// @summary Update a single review by its id.
// @description update a single review by its id.
// @tags update_review
// @accept json
// @produce json
// @success 200 {object} models.Review
// @router /application-inventory/review/:id [put]
// @param id path integer true "Review id"
// @param review body models.Review true "Review data"
func UpdateReview(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetReviewByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no review in db")
		return
	}

	updatedReview := models.Review{}
	err = c.BindJSON(&updatedReview)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateReview(database.DB, &updatedReview); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetReview(c)
}
