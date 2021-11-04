package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetStakeholderGroups godoc
// @summary Get all stakeholder groups.
// @description get all stakeholder groups.
// @tags get_stakeholder_groups
// @produce json
// @success 200 {object} models.Group
// @router /controls/stakeholder-group [get]
func GetStakeholderGroups(c *gin.Context) {
	stakeholderGroups, err := models.GetGroups(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stakeholderGroups)
}

// GetiStakeholderGroup godoc
// @summary Get a single stakeholder group by its id.
// @description get a single stakeholder group by its id.
// @tags get_stakeholder_group
// @produce json
// @success 200 {object} models.Group
// @router /controls/stakeholder-group/:id [get]
// @param id path integer true "Stakeholder Group id"
func GetStakeholderGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	stakeholderGroup, exists, err := models.GetGroupByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no stakeholder group in db")
		return
	}

	c.JSON(http.StatusOK, stakeholderGroup)
}

// CreateStakeholderGroup godoc
// @summary Create a single stakeholder group.
// @description create a single stakeholder group.
// @tags create_stakeholder_group
// @accept json
// @produce json
// @success 200 {object} models.Group
// @router /controls/stakeholder-group [post]
// @param stakeholder_group body models.Group true "Stakeholder Group data"
func CreateStakeholderGroup(c *gin.Context) {
	group := models.Group{}
	err := c.BindJSON(&group)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateGroup(database.DB, &group); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, group)
}

// DeleteStakeholderGroup godoc
// @summary Delete a single stakeholder group by its id.
// @description delete a single stakeholder group by its id.
// @tags delete_stakeholder_group
// @success 200 {object} models.Group
// @router /control/stakeholder-group/:id [delete]
// @param id path integer true "Stakeholder Group id"
func DeleteStakeholderGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetGroupByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no group in db")
		return
	}

	if err = models.DeleteGroup(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateStakeholderGroup godoc
// @summary Update a single stakeholder group by its id.
// @description update a single stakeholder group by its id.
// @tags update_stakeholder_group
// @accept json
// @produce json
// @success 200 {object} models.Group
// @router /controls/stakeholder-group/:id [put]
// @param id path integer true "Stakeholder Group id"
// @param stakeholder_group body models.Group true "Stakeholder Group data"
func UpdateStakeholderGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetGroupByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no group in db")
		return
	}

	updatedGroup := models.Group{}
	err = c.BindJSON(&updatedGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateGroup(database.DB, &updatedGroup); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetStakeholderGroup(c)
}
