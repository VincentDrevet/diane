package main

import (
	"github.com/VincentDrevet/configuration"
	"github.com/VincentDrevet/controllers"
	"github.com/VincentDrevet/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	configuration.ReadConfigurationFile("config.json")
	models.InitDB(configuration.AppConfig.DBPath)

	v1.GET("/servers", controllers.GetServers)
	v1.POST("/servers", controllers.AddServer)
	v1.GET("/servers/:id", controllers.GetServerByID)
	v1.GET("/tasks", controllers.GetTasks)
	v1.GET("/tasks/:id", controllers.GetTaskByID)
	v1.POST("/tasks", controllers.AddTask)
	v1.DELETE("/tasks/:id", controllers.DeleteTaskByID)

	router.Run()
}
