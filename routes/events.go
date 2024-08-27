package routes

import (
	"strconv"

	"example.com/restAPI/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	// gin.H is a shortcut for map[string]interface{}
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "Error Fetching All Events"})
		return
	}
	context.JSON(200, gin.H{"message": "Events Fetched!", "events": events})
}

func createEvent(context *gin.Context) {
	// extract data from the body
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not Extract Data"})
		return
	}

	// save the event
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Save failed"})
		return
	}

	context.JSON(201, gin.H{"message": "Event Created!", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Error Fetching Event"})
		return
	}

	context.JSON(200, gin.H{"message": "Event Found!", "event": event})
}
