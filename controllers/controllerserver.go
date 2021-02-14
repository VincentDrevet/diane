package controllers

import (
	"net/http"
	"strconv"

	"github.com/VincentDrevet/models"
	"github.com/gin-gonic/gin"
)

//GetServers return all servers in database
// [GET] /servers
func GetServers(c *gin.Context) {
	var servers []models.Server
	//models.DB.Find(&servers)
	models.DB.Preload("Tasks").Find(&servers)

	c.JSON(http.StatusOK, gin.H{
		"data": servers,
	})
}

//GetServerByID Return One server by ID
// [GET] /servers/:id
func GetServerByID(c *gin.Context) {
	var server models.Server

	if err := models.DB.Where("id = ?", c.Param("id")).Preload("Tasks").First(&server).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": server,
	})
}

//AddServer valid if body request is formatted as json, if not return badRequest if OK add in db server object and return created object.
// [POST] /servers
func AddServer(c *gin.Context) {
	var dtoserver models.DTOServer
	if err := c.ShouldBindJSON(&dtoserver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	server := models.Server{
		Name:       dtoserver.Name,
		Addr:       dtoserver.Addr,
		SSHPort:    dtoserver.SSHPort,
		User:       dtoserver.User,
		PrivateKey: dtoserver.PrivateKey,
		Tasks:      dtoserver.Tasks,
	}
	models.DB.Create(&server)

	c.JSON(http.StatusOK, gin.H{
		"data": server,
	})
}

//DeleteServerByID remove server and associate tasks
// [DELETE] /servers/:id
func DeleteServerByID(c *gin.Context) {
	var server models.Server
	models.DB.Delete(&server, c.Param("id"))
	models.DB.Where("server_id = ?", c.Param("id")).Delete(models.Task{})

	c.JSON(http.StatusOK, gin.H{})
}

//FullUpdateServer do full update of server object
// [PUT] /servers/:id
func FullUpdateServer(c *gin.Context) {
	var input models.Server
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

	//Update server
	err := models.DB.Debug().Model(&models.Server{}).Where("id = ?", c.Param("id")).Updates(models.Server{
		ID:         input.ID,
		Name:       input.Name,
		Addr:       input.Addr,
		SSHPort:    input.SSHPort,
		User:       input.User,
		PrivateKey: input.PrivateKey,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.DB.Where("server_id = ?", c.Param("id")).Delete(models.Task{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.DB.Create(input.Tasks)

	c.JSON(http.StatusNoContent, gin.H{})

}
