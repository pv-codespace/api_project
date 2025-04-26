package main

import (
	"net/http"
	"rest_api/db"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"data": events, "message": "success"})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse json data"})
		return
	}

	event.ID = "2"
	event.UserID = 2

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "created event", "data": event})
}
