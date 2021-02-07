package controllers

import (
	"net/http"

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
