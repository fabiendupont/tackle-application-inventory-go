package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-application-inventory-go/database"
	"github.com/fabiendupont/tackle-application-inventory-go/models"
	"net/http"
)

// GET /roleBindings
func GetRoleBindings(c *gin.Context) {
	roleBindings, err := models.GetRoleBindings(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, roleBindings)
}

// GET /roleBindings/:id
func GetRoleBinding(c *gin.Context) {
	id := c.Params.ByName("id")
	roleBinding, exists, err := models.GetRoleBindingByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no role binding in db")
		return
	}

	c.JSON(http.StatusOK, roleBinding)
}

// POST /roleBindings
func CreateRoleBinding(c *gin.Context) {
	roleBinding := models.RoleBinding{}
	err := c.BindJSON(&roleBinding)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateRoleBinding(database.DB, &roleBinding); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, roleBinding)
}

// DELETE /roleBindings/:id
func DeleteRoleBinding(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetRoleBindingByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no role binding in db")
		return
	}

	if err = models.DeleteRoleBinding(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}


func UpdateRoleBinding(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetRoleBindingByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no role binding in db")
		return
	}

	updatedRoleBinding := models.RoleBinding{}
	err = c.BindJSON(&updatedRoleBinding)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateRoleBinding(database.DB, &updatedRoleBinding); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetRoleBinding(c)
}
