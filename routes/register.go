package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	userID := c.GetInt64("userID")

	event, err := models.GetOneEvent(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	err = event.Register(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully"})

}
func cancelRegistration(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not retrieve event."})
		return
	}
	userID := c.GetInt64("userID")
	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration for event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cancelled registration for event successfully"})

}
