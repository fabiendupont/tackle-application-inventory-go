package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetOrganization godoc
// @summary Lists all the Organizations.
// @description get all the organizations
// @tags get_organizations
// @produce json
// @success 200 {object} []models.Organization
// @router /v2/organizations [get]
func GetOrganizations(c *gin.Context) {
	organizations, err := models.GetOrganizations(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organizations)
}

// GetOrganization godoc
// @summary Get a single organizcation by its id.
// @description get a single organization by its id.
// @tags get_organization
// @produce json
// @success 200 {object} models.Organization
// @router /v2/organizations/:id [get]
// @param id path integer true "Organization id"
func GetOrganization(c *gin.Context) {
	id := c.Params.ByName("id")
	organization, exists, err := models.GetOrganizationByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no organization in db")
		return
	}

	c.JSON(http.StatusOK, organization)
}

// CreateOrganization godoc
// @summary Create an organization
// @description create an organization
// @tags create_organization
// @accept json
// @produce json
// @success 200 {object} models.Organization
// @router /v2/organizations [post]
// @param application body models.Organization true "Organization data"
func CreateOrganization(c *gin.Context) {
	organization := models.Organization{}
	err := c.BindJSON(&organization)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateOrganization(database.DB, &organization); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organization)
}

// DeleteOrganization godoc
// @summary Delete an organization
// @description delete an organization
// @tags delete_organization
// @success 200 {object} models.Organization
// @router /v2/organizations/:id [delete]
// @param id path integer true "Organization id"
func DeleteOrganization(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetOrganizationByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no organization in db")
		return
	}

	if err = models.DeleteOrganization(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateOrganization godoc
// @summary Update an organization
// @description update an organization
// @tags update_organization
// @accept json
// @produce json
// @success 200 {object} models.Organization
// @router /v2/organizations/:id [put]
// @param id path integer true "Organization id"
// @param application body models.Organization true "Organization data"
func UpdateOrganization(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetOrganizationByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no organization in db")
		return
	}

	updatedOrganization := models.Organization{}
	err = c.BindJSON(&updatedOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateOrganization(database.DB, &updatedOrganization); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetOrganization(c)
}
