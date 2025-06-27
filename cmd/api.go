package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willnee/cloud-services/pkg/db"
	"github.com/willnee/cloud-services/pkg/handlers"
	"log"
)

func main() {
	db.Connect()
	handlers.Init()
	r := gin.Default()

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTaskByID)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
