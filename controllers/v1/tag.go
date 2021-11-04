package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetTags godoc
// @summary Get all tags.
// @description get all tags.
// @tags get_tags
// @produce json
// @success 200 {object} models.Tag
// @router /controls/tag [get]
func GetTags(c *gin.Context) {
	tags, err := models.GetTags(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetTag godoc
// @summary Get a single tag by its id.
// @description get a single tag by its id.
// @tags get_tag
// @produce json
// @success 200 {object} models.Tag
// @router /controls/tag/:id [get]
// @param id path integer true "Tag id"
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

// CreateTag godoc
// @summary Create a single tag.
// @description create a single tag.
// @tags create_tag
// @accept json
// @produce json
// @success 200 {object} models.Tag
// @router /controls/tag [post]
// @param tag body models.Tag true "Tag data"
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

// Delete godoc
// @summary Delete a single tag by its id.
// @description delete a single tag by its id.
// @tags delete_tag
// @success 200 {object} models.Tag
// @router /controls/tag/:id [delete]
// @param id path integer true "Tag id"
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

// UpdateTag godoc
// @summary Update a single tag by its id.
// @description update a single tag by its id.
// @tags update_tag
// @accept json
// @produce json
// @success 200 {object} models.Tag
// @router /controls/tag/:id [put]
// @param id path integer true "Tag id"
// @param tag body models.Tag true "Tag data"
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
