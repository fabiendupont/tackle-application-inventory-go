package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetApplications godoc
// @summary Lists all the applications.
// @description get all the application
// @tags get_applications
// @produce json
// @success 200 {object} []models.Application
// @router /application-inventory/application [get]
func GetApplications(c *gin.Context) {
	applications, err := models.GetApplications(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, applications)
}

// GetApplication godoc
// @summary Get a single application by its id.
// @description get a single application by its id.
// @tags get_application
// @produce json
// @success 200 {object} models.Application
// @router /application-inventory/application/:id [get]
// @param id path integer true "Application id"
func GetApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	application, exists, err := models.GetApplicationByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no application in db")
		return
	}

	c.JSON(http.StatusOK, application)
}

// CreateApplication godoc
// @summary Create an application
// @description create an application
// @tags create_application
// @accept json
// @produce json
// @success 200 {object} models.Application
// @router /application-inventory/application [post]
// @param application body models.Application true "Application data"
func CreateApplication(c *gin.Context) {
	application := models.Application{}
	err := c.BindJSON(&application)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateApplication(database.DB, &application); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, application)
}

// DeleteApplication godoc
// @summary Delete an application
// @description delete an application
// @tags delete_application
// @success 200 {object} models.Application
// @router /application-inventory/application/:id [delete]
// @param id path integer true "Application id"
func DeleteApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetApplicationByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no application in db")
		return
	}

	if err = models.DeleteApplication(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateApplication godoc
// @summary Update an application
// @description update an application
// @tags update_application
// @accept json
// @produce json
// @success 200 {object} models.Application
// @router /application-inventory/application/:id [put]
// @param id path integer true "Application id"
// @param application body models.Application true "Application data"
func UpdateApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetApplicationByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no application in db")
		return
	}

	updatedApplication := models.Application{}
	err = c.BindJSON(&updatedApplication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateApplication(database.DB, &updatedApplication); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetApplication(c)
}
