package main

import (
	"interview/controllers/activitycontroller"
	"interview/controllers/todocontroller"
	"interview/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// activityroute
	r.GET("/activity-groups", activitycontrollers.Index)
	r.GET("/activity-groups/:id", activitycontrollers.Show)
	r.POST("/activity-groups", activitycontrollers.Create)
	r.PATCH("/activity-groups/:id", activitycontrollers.Update)
	r.DELETE("/activity-groups/:id", activitycontrollers.Delete)

	// todoroute
	r.GET("/todo-items", todocontrollers.Index)
	r.GET("/todo-items/:id", todocontrollers.Show)
	r.POST("/todo-items", todocontrollers.Create)
	r.PATCH("/todo-items/:id", todocontrollers.Update)
	r.DELETE("/todo-items/:activity_id", todocontrollers.Delete)
	r.Run(":8000")
}
