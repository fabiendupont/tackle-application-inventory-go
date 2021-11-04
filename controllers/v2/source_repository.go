package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GET /source_repositories
func GetSourceRepositories(c *gin.Context) {
	source_repositories, err := models.GetSourceRepositories(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, source_repositories)
}

// GET /source_repositories/:id
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

// POST /source_repositories
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

// DELETE /source_repositories/:id
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

// PUT /source_repositories/:id
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
