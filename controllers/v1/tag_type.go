package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetTagTypes godoc
// @summary Get all tag types.
// @description get all tag types.
// @tags get_tag_types
// @produce json
// @success 200 {object} models.TagType
// @router /controls/tag-type [get]
func GetTagTypes(c *gin.Context) {
	tagTypes, err := models.GetTagTypes(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tagTypes)
}

// GetTagType godoc
// @summary Get a single tag type by its id.
// @description get a single tag type by its id.
// @tags get_tag_type
// @produce json
// @success 200 {object} models.TagType
// @router /controls/tag-type/:id [get]
// @param id path integer true "Tag Type id"
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

// CreateTagType godoc
// @summary Create a single tag type.
// @description create a single tag type.
// @tags create_tag_type
// @accept json
// @produce json
// @success 200 {object} models.TagType
// @router /controls/tag-type [post]
// @param tag_type body models.TagType true "Tag Type data"
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

// DeleteTagType godoc
// @summary Delete a single tag type by its id.
// @description delete a single tag type by its id.
// @tags delete_tag_type
// @success 200 {object} models.TagType
// @router /controls/tag-type/:id [delete]
// @param id path integer true "Tag Type id"
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

// UpdateTagType godoc
// @summary Update a single tag type by its id.
// @description update a single tag type by its id.
// @tags update_tag_type
// @accept json
// @produce json
// @success 200 {object} models.TagType
// @router /controls/tag-type/:id [put]
// @param id path integer true "Tag Type id"
// @param tag_type body models.TagType true "Tag Type data"
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
