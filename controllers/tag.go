package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/fabiendupont/tackle-application-inventory-go/models"
	"net/http"
)

// GET /tags
func GetTags(c *gin.Context) {
	tags, err := models.GetTags(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GET /tags/:id
func GetTag(c *gin.Context) {
	id := c.Params.ByName("id")
	tag, exists, err := models.GetTagByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no tag in db")
		return
	}

	c.JSON(http.StatusOK, tag)
}

// POST /tags
func CreateTag(c *gin.Context) {
	tag := models.Tag{}
	err := c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateTag(database.DB, &tag); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tag)
}

// DELETE /tags/:id
func DeleteTag(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetTagByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no tag in db")
		return
	}

	if err = models.DeleteTag(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}


func UpdateTag(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetTagByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no tag in db")
		return
	}

	updatedTag := models.Tag{}
	err = c.BindJSON(&updatedTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateTag(database.DB, &updatedTag); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetTag(c)
}
