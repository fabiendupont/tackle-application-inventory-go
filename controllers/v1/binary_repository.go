package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetBinaryRepositories godoc
// @summary Get all binary repositories.
// @description get all binary repositories.
// @tags get_binary_repositories
// @produce json
// @success 200 {object} models.BinaryRepository
// @router /application-inventory/binary-repository [get]
func GetBinaryRepositories(c *gin.Context) {
	binary_repositories, err := models.GetBinaryRepositories(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, binary_repositories)
}

// GetBinaryRepository godoc
// @summary Get a single binary repository by its id.
// @description get a single binary repository by its id.
// @tags get_binary_repository
// @produce json
// @success 200 {object} models.BinaryRepository
// @router /application-inventory/binary-repository/:id [get]
// @param id path integer true "Binary Repsoritory id"
func GetBinaryRepository(c *gin.Context) {
	id := c.Params.ByName("id")
	binary_repository, exists, err := models.GetBinaryRepositoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no binary_repository in db")
		return
	}

	c.JSON(http.StatusOK, binary_repository)
}

// CreateBinaryRepository godoc
// @summary Create a single binary repository.
// @description create a single binary repository.
// @tags get_binary_repository
// @accept json
// @produce json
// @success 200 {object} models.BinaryRepository
// @router /application-inventory/binary-repository [post]
// @param binary_repository body models.BinaryRepository true "Binary Repository data"
func CreateBinaryRepository(c *gin.Context) {
	binary_repository := models.BinaryRepository{}
	err := c.BindJSON(&binary_repository)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateBinaryRepository(database.DB, &binary_repository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, binary_repository)
}

// DeleteBinaryRepository godoc
// @summary Delete a single binary repository by its id.
// @description delete a single binary repository by its id.
// @tags delete_binary_repository
// @success 200 {object} models.BinaryRepository
// @router /application-inventory/binary-repository/:id [delete]
// @param id path integer true "Binary Repository id"
func DeleteBinaryRepository(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetBinaryRepositoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no binary_repository in db")
		return
	}

	if err = models.DeleteBinaryRepository(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateBinaryRepository godoc
// @summary Update a single binary repository by its id.
// @description update a single binary repository by its id.
// @tags update_binary_repository
// @accept json
// @produce json
// @success 200 {object} models.BinaryRepository
// @router /application-inventory/binary_repository/:id [put]
// @param id path integer true "Binary Repository id"
// @param binary_repository body models.BinaryRepository true "Binary Repository data"
func UpdateBinaryRepository(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetBinaryRepositoryByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no binary_repository in db")
		return
	}

	updatedBinaryRepository := models.BinaryRepository{}
	err = c.BindJSON(&updatedBinaryRepository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateBinaryRepository(database.DB, &updatedBinaryRepository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetBinaryRepository(c)
}
