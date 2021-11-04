package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetBusinessServices godoc
// @summary Get all business services.
// @description get all business services.
// @tags get_business_services
// @produce json
// @success 200 {object} []models.BusinessService
// @router /controls/business-service [get]
func GetBusinessServices(c *gin.Context) {
	businessServices, err := models.GetBusinessServices(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, businessServices)
}

// GetBusinessService godoc
// @summary Get a single business service by its id.
// @description get a single business service by its id.
// @tags get_business_service
// @produce json
// @success 200 {object} models.BusinessService
// @router /controls/business-service/:id [get]
// @param id path integer true "Business Service id"
func GetBusinessService(c *gin.Context) {
	id := c.Params.ByName("id")
	businessService, exists, err := models.GetBusinessServiceByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no business service in db")
		return
	}

	c.JSON(http.StatusOK, businessService)
}

// CreateBusinessService godoc
// @summary Create a business service.
// @description create a business service.
// @tags create_business_service
// @accept json
// @produce json
// @success 200 {object} models.BusinessService
// @router /controls/business-service [post]
// @param business_service body models.BusinessService true "Business Service data"
func CreateBusinessService(c *gin.Context) {
	businessService := models.BusinessService{}
	err := c.BindJSON(&businessService)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateBusinessService(database.DB, &businessService); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, businessService)
}

// DeleteBusinessService godoc
// @summary Delete a single business service by its id.
// @description delete a single business service by its id.
// @tags delete_business_service
// @success 200 {object} models.BusinessService
// @router /controls/business-service/:id [delete]
// @param id path integer true "Business Service id"
func DeleteBusinessService(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetBusinessServiceByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no business service in db")
		return
	}

	if err = models.DeleteBusinessService(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateBusinessService godoc
// @summary Update a single business service by its id.
// @description update a single business_service by its id.
// @tags update_business_service
// @accept json
// @produce json
// @success 200 {object} models.BusinessService
// @router /controls/business-service/:id [put]
// @param id path integer true "Business Service id"
// @param business_service body models.BusinessService true "Business Service data"
func UpdateBusinessService(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetBusinessServiceByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no business service in db")
		return
	}

	updatedBusinessService := models.BusinessService{}
	err = c.BindJSON(&updatedBusinessService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateBusinessService(database.DB, &updatedBusinessService); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetBusinessService(c)
}
