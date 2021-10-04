package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/fabiendupont/tackle-application-inventory-go/models"
	"net/http"
)

// GET /groups
func GetGroups(c *gin.Context) {
	groups, err := models.GetGroups(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GET /groups/:id
func GetGroup(c *gin.Context) {
	id := c.Params.ByName("id")
	group, exists, err := models.GetGroupByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no group in db")
		return
	}

	c.JSON(http.StatusOK, group)
}

// POST /groups
func CreateGroup(c *gin.Context) {
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

// DELETE /groups/:id
func DeleteGroup(c *gin.Context) {
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


func UpdateGroup(c *gin.Context) {
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

	GetGroup(c)
}
