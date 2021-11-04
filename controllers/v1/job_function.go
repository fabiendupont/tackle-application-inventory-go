package controllers

import (
	"github.com/gin-gonic/gin"
        "github.com/fabiendupont/tackle-hub/database"
	"github.com/fabiendupont/tackle-hub/models"
	"net/http"
)

// GetJobFunctions godoc
// @summary Get all job functions.
// @description get all job functions.
// @tags get_job_functions
// @produce json
// @success 200 {object} models.JobFunction
// @router /controls/job-function [get]
func GetJobFunctions(c *gin.Context) {
	job_functions, err := models.GetJobFunctions(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, job_functions)
}

// GetJobFunction godoc
// @summary Get a single job function by its id.
// @description get a single job function by its id.
// @tags get_job_function
// @produce json
// @success 200 {object} models.JobFunction
// @router /controls/job-function/:id [get]
// @param id path integer true "Job Function id"
func GetJobFunction(c *gin.Context) {
	id := c.Params.ByName("id")
	job_function, exists, err := models.GetJobFunctionByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no job_function in db")
		return
	}

	c.JSON(http.StatusOK, job_function)
}

// CreateApplication godoc
// @summary Create a single job function.
// @description create a single job function.
// @tags create_job_function
// @accept json
// @produce json
// @success 200 {object} models.JobFunction
// @router /controls/job-function [post]
// @param application body models.JobFunction true "Job Function data"
func CreateJobFunction(c *gin.Context) {
	job_function := models.JobFunction{}
	err := c.BindJSON(&job_function)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.CreateJobFunction(database.DB, &job_function); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, job_function)
}

// DeleteJobFunction godoc
// @summary Delete a single job function by its id.
// @description delete a single job function by its id.
// @tags delete_job_function
// @success 200 {object} models.JobFunction
// @router /controls/job-function/:id [delete]
// @param id path integer true "Job Function id"
func DeleteJobFunction(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetJobFunctionByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no job_function in db")
		return
	}

	if err = models.DeleteJobFunction(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UpdateJobFunction godoc
// @summary Update a single job function by its id.
// @description update a single job function by its id.
// @tags update_job_function
// @accept json
// @produce json
// @success 200 {object} models.JobFunction
// @router /controls/job-function/:id [put]
// @param id path integer true "Job Function id"
// @param application body models.JobFunction true "Job Function data"
func UpdateJobFunction(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := models.GetJobFunctionByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no job_function in db")
		return
	}

	updatedJobFunction := models.JobFunction{}
	err = c.BindJSON(&updatedJobFunction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = models.UpdateJobFunction(database.DB, &updatedJobFunction); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	GetJobFunction(c)
}
