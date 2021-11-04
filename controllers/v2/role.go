package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GET /roles
func GetRoles(c *gin.Context) {
	roles, err := models.GetRoles(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, roles)
}

// GET /roles/:id
func GetRole(c *gin.Context) {
	id := c.Params.ByName("id")
	role, exists, err := models.GetRoleByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no role in db")
		return
	}

	c.JSON(http.StatusOK, role)
}

// POST /roles
func CreateRole(c *gin.Context) {
	role := models.Role{}
	err := c.BindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateRole(database.DB, &role); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, role)
}

// DELETE /roles/:id
func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetRoleByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no role in db")
		return
	}

	if err = models.DeleteRole(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}


func UpdateRole(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetRoleByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no role in db")
		return
	}

	updatedRole := models.Role{}
	err = c.BindJSON(&updatedRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateRole(database.DB, &updatedRole); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetRole(c)
}
