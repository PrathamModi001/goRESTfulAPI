package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// Bind the JSON body to the user struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, gin.H{"message": "User created successfully"})

}