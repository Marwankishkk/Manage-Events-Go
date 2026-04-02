package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events."})
		return
	}

	c.JSON(http.StatusOK, events)
}
func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	event, err := models.GetOneEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	c.JSON(http.StatusOK, event)
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
