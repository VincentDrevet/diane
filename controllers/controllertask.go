package controllers

import (
	"net/http"
	"strconv"

	"github.com/VincentDrevet/models"
	"github.com/gin-gonic/gin"
)

//GetTasks return all tasks from databases
//[GET] /tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

//GetTaskByID return task by ID from DB
//[GET] /tasks/:id
func GetTaskByID(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

//AddTask Add new task via DTOTask struct from DTO struct
//[POST] /tasks
func AddTask(c *gin.Context) {
	var dtotask models.DTOTask
	if err := c.ShouldBindJSON(&dtotask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := models.Task{
		Hour:        dtotask.Hour,
		Minute:      dtotask.Minute,
		Second:      dtotask.Second,
		Day:         dtotask.Day,
		Periodicity: dtotask.Periodicity,
		ServerID:    dtotask.ServerID,
	}

	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})

}

//FullUpdateTask do full update of task
//[PUT] /tasks/:id
func FullUpdateTask(c *gin.Context) {

	//var task models.Task
	var input models.Task
	/*


		if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	*/

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if strconv.Itoa(int(input.ID)) != c.Param("id") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID mismatch",
		})
		return
	}

	err := models.DB.Model(&models.Task{}).Where("id = ?", c.Param("id")).Updates(models.Task{
		ID:          input.ID,
		Hour:        input.Hour,
		Minute:      input.Minute,
		Second:      input.Second,
		Day:         input.Day,
		Periodicity: input.Day,
		ServerID:    input.ServerID,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}

//DeleteTaskByID delete task by ID from database
//[DELETE] /tasks/:id
func DeleteTaskByID(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
