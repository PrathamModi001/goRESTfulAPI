package main

import (
	"example.com/restAPI/db"
	"example.com/restAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()
	// Create a new Gin instance
	server := gin.Default()

	// Register the routes
	routes.RegisterRoutes(server)

	// mention the port
	server.Run(":8080")
}
