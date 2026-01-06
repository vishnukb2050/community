package handlers

import (
	"calendar-service/config"
	"calendar-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "calendar-service",
		"time":    time.Now(),
	})
}

// GetCalendar - Get all events for a user (unified calendar)
func GetCalendar(c *gin.Context) {
	userID := c.GetString("user_id")

	var events []models.CalendarEvent
	if err := config.DB.Where("user_id = ?", userID).
		Order("event_date ASC, event_time ASC").
		Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch calendar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

// GetEventsByDate - Get events for specific date
func GetEventsByDate(c *gin.Context) {
	userID := c.GetString("user_id")
	date := c.Param("date") // YYYY-MM-DD format

	var events []models.CalendarEvent
	if err := config.DB.Where("user_id = ? AND DATE(event_date) = ?", userID, date).
		Order("event_time ASC").
		Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events, "date": date})
}

// CreateEvent - Create new calendar event
func CreateEvent(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse date
	eventDate, err := time.Parse("2006-01-02", req.EventDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// Parse time if provided
	var eventTime time.Time
	if req.EventTime != "" {
		eventTime, err = time.Parse("15:04", req.EventTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format, use HH:MM"})
			return
		}
	}

	userUUID, _ := uuid.Parse(userID)
	event := models.CalendarEvent{
		UserID:      userUUID,
		Title:       req.Title,
		Description: req.Description,
		EventDate:   eventDate,
		EventTime:   eventTime,
		EventType:   "personal",
	}

	if err := config.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

// UpdateEvent - Update existing event
func UpdateEvent(c *gin.Context) {
	userID := c.GetString("user_id")
	eventID := c.Param("id")

	var event models.CalendarEvent
	if err := config.DB.Where("id = ? AND user_id = ?", eventID, userID).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if req.Title != "" {
		event.Title = req.Title
	}
	if req.Description != "" {
		event.Description = req.Description
	}
	if req.EventDate != "" {
		eventDate, err := time.Parse("2006-01-02", req.EventDate)
		if err == nil {
			event.EventDate = eventDate
		}
	}

	config.DB.Save(&event)

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

// DeleteEvent - Delete event
func DeleteEvent(c *gin.Context) {
	userID := c.GetString("user_id")
	eventID := c.Param("id")

	result := config.DB.Where("id = ? AND user_id = ?", eventID, userID).
		Delete(&models.CalendarEvent{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
