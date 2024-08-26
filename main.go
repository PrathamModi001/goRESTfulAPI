package main
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin instance
	server := gin.Default()

	server.GET("/events", getEvents)

	// mention the port
	server.Run(":8080")
}

// 
func getEvents(context *gin.Context){
	// gin.H is a shortcut for map[string]interface{}
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}