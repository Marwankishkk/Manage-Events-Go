package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	// This is a placeholder for the main function.
	// In a real application, you would initialize your server, database connections, etc. here.
	server.Run(":8080") // Start the server on port 8080
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events."})
		return
	}

	c.JSON(http.StatusOK, events)
}
func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully", "event": event})
}
