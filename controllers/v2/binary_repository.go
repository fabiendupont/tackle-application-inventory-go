package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GET /binary_repositories
func GetBinaryRepositories(c *gin.Context) {
	binary_repositories, err := models.GetBinaryRepositories(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, binary_repositories)
}

// GET /binary_repositories/:id
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

// POST /binary_repositories
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

// DELETE /binary_repositories/:id
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

// PUT /binary_repositories/:id
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
