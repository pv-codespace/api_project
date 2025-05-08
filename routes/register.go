package routes

import (
	"net/http"
	"rest_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 24)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Faied to get event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Faied to register"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully register for event"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 24)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Faied to get event"})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Faied to deregister"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully Deregister from event"})
}
