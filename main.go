package main
import (
	"github.com/gin-gonic/gin"
	"example.com/restAPI/models"
)

func main() {
	// Create a new Gin instance
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// mention the port
	server.Run(":8080")
}

func getEvents(context *gin.Context){
	// gin.H is a shortcut for map[string]interface{}
	events := models.GetAllEvents()
	context.JSON(200, gin.H{"message": "Event Created!", "events": events})
}

func createEvent(context *gin.Context) {
	// extract data from the body
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// save the event
	event.ID = 1
	event.UserID = 1

	context.JSON(201, gin.H{"message": "Event Created!", "event": event})
}