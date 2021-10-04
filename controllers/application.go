package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/fabiendupont/tackle-application-inventory-go/models"
	"net/http"
)

// GET /applications
func GetApplications(c *gin.Context) {
	applications, err := models.GetApplications(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, applications)
}

// GET /applications/:id
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

// POST /applications
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

// DELETE /applications/:id
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
