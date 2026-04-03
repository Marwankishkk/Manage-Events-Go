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
	userID := c.GetInt64("userID")
	event.UserID = userID
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully", "event": event})
}
func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	userID := c.GetInt64("userID")
	event, err := models.GetOneEvent(id)
	if event.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update event"})
		return

	}
	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEvent.ID = id
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event."})
		return
	}
	updatedEvent.UserID = 1
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})

}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	userID := c.GetInt64("userID")
	event, err := models.GetOneEvent(id)

	if userID != event.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete event"})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
