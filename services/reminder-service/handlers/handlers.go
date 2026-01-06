package handlers

import (
	"net/http"
	"reminder-service/config"
	"reminder-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "reminder-service",
		"time":    time.Now(),
	})
}

// GetReminders - List all user reminders with optional filter
func GetReminders(c *gin.Context) {
	userID := c.GetString("user_id")
	filter := c.Query("filter") // upcoming, overdue, paid

	query := config.DB.Where("user_id = ?", userID)

	if filter == "upcoming" {
		query = query.Where("status = ? AND due_date >= ?", "pending", time.Now())
	} else if filter == "overdue" {
		query = query.Where("status = ? AND due_date < ?", "pending", time.Now())
	} else if filter == "paid" {
		query = query.Where("status = ?", "paid")
	}

	var reminders []models.Reminder
	if err := query.Order("due_date ASC").Find(&reminders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reminders": reminders, "count": len(reminders)})
}

// CreateReminder - Create new reminder
func CreateReminder(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateReminderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	userUUID, _ := uuid.Parse(userID)
	reminder := models.Reminder{
		UserID:         userUUID,
		Title:          req.Title,
		Amount:         req.Amount,
		DueDate:        dueDate,
		IsRecurring:    req.IsRecurring,
		RecurrenceType: req.RecurrenceType,
		Status:         "pending",
	}

	if err := config.DB.Create(&reminder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reminder"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Reminder created successfully",
		"reminder": reminder,
	})
}

// UpdateReminder - Update existing reminder
func UpdateReminder(c *gin.Context) {
	userID := c.GetString("user_id")
	reminderID := c.Param("id")

	var reminder models.Reminder
	if err := config.DB.Where("id = ? AND user_id = ?", reminderID, userID).First(&reminder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	var req models.UpdateReminderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Title != "" {
		reminder.Title = req.Title
	}
	if req.Amount > 0 {
		reminder.Amount = req.Amount
	}
	if req.Status != "" {
		reminder.Status = req.Status
	}
	if req.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", req.DueDate)
		if err == nil {
			reminder.DueDate = dueDate
		}
	}

	config.DB.Save(&reminder)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Reminder updated successfully",
		"reminder": reminder,
	})
}

// DeleteReminder - Delete reminder
func DeleteReminder(c *gin.Context) {
	userID := c.GetString("user_id")
	reminderID := c.Param("id")

	result := config.DB.Where("id = ? AND user_id = ?", reminderID, userID).
		Delete(&models.Reminder{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted successfully"})
}

// MarkPaid - Mark reminder as paid
func MarkPaid(c *gin.Context) {
	userID := c.GetString("user_id")
	reminderID := c.Param("id")

	var reminder models.Reminder
	if err := config.DB.Where("id = ? AND user_id = ?", reminderID, userID).First(&reminder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	reminder.Status = "paid"
	config.DB.Save(&reminder)

	// If recurring, create next reminder
	if reminder.IsRecurring {
		var nextDueDate time.Time
		if reminder.RecurrenceType == "monthly" {
			nextDueDate = reminder.DueDate.AddDate(0, 1, 0)
		} else if reminder.RecurrenceType == "yearly" {
			nextDueDate = reminder.DueDate.AddDate(1, 0, 0)
		}

		if !nextDueDate.IsZero() {
			nextReminder := models.Reminder{
				UserID:         reminder.UserID,
				Title:          reminder.Title,
				Amount:         reminder.Amount,
				DueDate:        nextDueDate,
				IsRecurring:    true,
				RecurrenceType: reminder.RecurrenceType,
				Status:         "pending",
			}
			config.DB.Create(&nextReminder)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Reminder marked as paid",
		"reminder": reminder,
	})
}

// SnoozeReminder - Snooze reminder for later
func SnoozeReminder(c *gin.Context) {
	userID := c.GetString("user_id")
	reminderID := c.Param("id")

	var req struct {
		Days int `json:"days" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Days required"})
		return
	}

	var reminder models.Reminder
	if err := config.DB.Where("id = ? AND user_id = ?", reminderID, userID).First(&reminder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	reminder.Status = "snoozed"
	reminder.SnoozedUntil = time.Now().AddDate(0, 0, req.Days)
	config.DB.Save(&reminder)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Reminder snoozed",
		"snoozed_until": reminder.SnoozedUntil,
	})
}
