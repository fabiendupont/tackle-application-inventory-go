package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetSourceRepositories godoc
// @summary Get all source repositories.
// @description get all source repositories.
// @tags get_source_repositories
// @produce json
// @success 200 {object} models.SourceRepository
// @router /application-inventory/source-repository [get]
func GetSourceRepositories(c *gin.Context) {
	source_repositories, err := models.GetSourceRepositories(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, source_repositories)
}

// GetSourceRepository godoc
// @summary Get a single source repository by its id.
// @description get a single source repository by its id.
// @tags get_source_repository
// @produce json
// @success 200 {object} models.SourceRepository
// @router /application-inventory/source-repository/:id [get]
// @param id path integer true "Source Repsoritory id"
func GetSourceRepository(c *gin.Context) {
	id := c.Params.ByName("id")
	source_repository, exists, err := models.GetSourceRepositoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no source_repository in db")
		return
	}

	c.JSON(http.StatusOK, source_repository)
}

// CreateSourceRepository godoc
// @summary Create a single source repository.
// @description create a single source repository.
// @tags get_source_repository
// @accept json
// @produce json
// @success 200 {object} models.SourceRepository
// @router /application-inventory/source-repository [post]
// @param source_repository body models.SourceRepository true "Source Repository data"
func CreateSourceRepository(c *gin.Context) {
	source_repository := models.SourceRepository{}
	err := c.BindJSON(&source_repository)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateSourceRepository(database.DB, &source_repository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, source_repository)
}

// DeleteSourceRepository godoc
// @summary Delete a single source repository by its id.
// @description delete a single source repository by its id.
// @tags delete_source_repository
// @success 200 {object} models.SourceRepository
// @router /application-inventory/source-repository/:id [delete]
// @param id path integer true "Source Repository id"
func DeleteSourceRepository(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetSourceRepositoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no source_repository in db")
		return
	}

	if err = models.DeleteSourceRepository(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateSourceRepository godoc
// @summary Update a single source repository by its id.
// @description update a single source repository by its id.
// @tags update_source_repository
// @accept json
// @produce json
// @success 200 {object} models.SourceRepository
// @router /application-inventory/source_repository/:id [put]
// @param id path integer true "Source Repository id"
// @param source_repository body models.SourceRepository true "Source Repository data"
func UpdateSourceRepository(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetSourceRepositoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no source_repository in db")
		return
	}

	updatedSourceRepository := models.SourceRepository{}
	err = c.BindJSON(&updatedSourceRepository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateSourceRepository(database.DB, &updatedSourceRepository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetSourceRepository(c)
}
