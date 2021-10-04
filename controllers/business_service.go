package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/fabiendupont/tackle-application-inventory-go/models"
	"net/http"
)

// GET /business_services
func GetBusinessServices(c *gin.Context) {
	businessServices, err := models.GetBusinessServices(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, businessServices)
}

// GET /business_services/:id
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

// POST /business_services
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

// DELETE /business_services/:id
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
