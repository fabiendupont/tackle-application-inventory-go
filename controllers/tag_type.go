package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/fabiendupont/tackle-application-inventory-go/models"
	"net/http"
)

// GET /tagTypes
func GetTagTypes(c *gin.Context) {
	tagTypes, err := models.GetTagTypes(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tagTypes)
}

// GET /tagTypes/:id
func GetTagType(c *gin.Context) {
	id := c.Params.ByName("id")
	tagType, exists, err := models.GetTagTypeByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no tag type in db")
		return
	}

	c.JSON(http.StatusOK, tagType)
}

// POST /tagTypes
func CreateTagType(c *gin.Context) {
	tagType := models.TagType{}
	err := c.BindJSON(&tagType)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateTagType(database.DB, &tagType); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tagType)
}

// DELETE /tagTypes/:id
func DeleteTagType(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetTagTypeByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no tag type in db")
		return
	}

	if err = models.DeleteTagType(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}


func UpdateTagType(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetTagTypeByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no tag type in db")
		return
	}

	updatedTagType := models.TagType{}
	err = c.BindJSON(&updatedTagType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateTagType(database.DB, &updatedTagType); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetTagType(c)
}
