package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

type Stakeholder struct {
	ID		uint	`json:"id"`
	Email		string	`json:"email" binding:"required"`
	DisplayName	string	`json:"display_name" binding:"required"`
	JobFunctionID	uint	`json:"job_function_id" binding:"required"`
}

func UserToStakeholder(user models.User) Stakeholder {
	return Stakeholder{
		user.ID,
		user.Email,
		user.DisplayName,
		user.JobFunctionID,
	}
}

func StakeholderToUser(stakeholder Stakeholder) (user models.User) {
	user.Email = stakeholder.Email
	user.DisplayName = stakeholder.DisplayName
	user.JobFunctionID = stakeholder.JobFunctionID
	return user
}

// GetStakeholders godoc
// @summary Get all stakeholders.
// @description get all stakeholders.
// @tags get_stakeholders
// @produce json
// @success 200 {object} Stakeholder
// @router /controls/stakeholder [get]
func GetStakeholders(c *gin.Context) {
	users, err := models.GetUsers(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	stakeholders := []Stakeholder{}
	for _, u := range(users) {
		stakeholders = append(stakeholders, UserToStakeholder(u))
	}

	c.JSON(http.StatusOK, stakeholders)
}

// GetStakeholder godoc
// @summary Get a single stakeholder by its id.
// @description get a single stakeholder by its id.
// @tags get_stakeholder
// @produce json
// @success 200 {object} Stakeholder
// @router /controls/stakeholder/:id [get]
// @param id path integer true "Stakeholder id"
func GetStakeholder(c *gin.Context) {
	id := c.Params.ByName("id")
	user, exists, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	stakeholder := UserToStakeholder(user)

	c.JSON(http.StatusOK, stakeholder)
}

// CreateStakeholder godoc
// @summary Create a single stakeholder.
// @description create a single stakeholder.
// @tags create stakeholder
// @accept json
// @produce json
// @success 200 {object} Stakeholder
// @router /controls/stakeholder [post]
// @param stakeholder body Stakeholder true "Stakeholder data"
func CreateStakeholder(c *gin.Context) {
	stakeholder := Stakeholder{}
	err := c.BindJSON(&stakeholder)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := StakeholderToUser(stakeholder)

	if err := models.CreateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	createdStakeholder := UserToStakeholder(user)
	c.JSON(http.StatusOK, createdStakeholder)
}

// DeleteStakeholder godoc
// @summary Delete a single stakeholder by its id.
// @description delete a single stakeholder by its id.
// @tags delete_stakeholder
// @success 200 {object} Stakeholder
// @router /controls/stakeholder/:id [delete]
// @param id path integer true "Stakeholder id"
func DeleteStakeholder(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	if err = models.DeleteUser(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateStakeholder godoc
// @summary Update a single stakeholder by its id.
// @description update a single stakeholder by its id.
// @tags update_stakeholder
// @accept json
// @produce json
// @success 200 {object} Stakeholder
// @router /controls/stakeholder/:id [put]
// @param id path integer true "Stakeholder id"
// @param stakeholder body Stakeholder true "Stakeholder data"
func UpdateStakeholder(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	updatedStakeholder := Stakeholder{}
	err = c.BindJSON(&updatedStakeholder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	updatedUser := StakeholderToUser(updatedStakeholder)

	if err = models.UpdateUser(database.DB, &updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetStakeholder(c)
}
