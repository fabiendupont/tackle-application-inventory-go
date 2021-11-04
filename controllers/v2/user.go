package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GET /users
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, exists, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	c.JSON(http.StatusOK, user)
}

// POST /users
func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	if err = models.DeleteUser(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}


func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetUserByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no user in db")
		return
	}

	updatedUser := models.User{}
	err = c.BindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateUser(database.DB, &updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetUser(c)
}
