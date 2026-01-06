package handlers

import (
	"net/http"
	"notes-service/config"
	"notes-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "notes-service",
		"time":    time.Now(),
	})
}

// GetNotes - List all user notes
func GetNotes(c *gin.Context) {
	userID := c.GetString("user_id")

	var notes []models.Note
	if err := config.DB.Where("user_id = ?", userID).
		Order("updated_at DESC").
		Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

// CreateNote - Create new note
func CreateNote(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUUID, _ := uuid.Parse(userID)
	note := models.Note{
		UserID:   userUUID,
		Title:    req.Title,
		Content:  req.Content,
		NoteType: req.NoteType,
	}

	if note.NoteType == "" {
		note.NoteType = "note"
	}

	if err := config.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Note created successfully",
		"note":    note,
	})
}

// GetNote - Get single note by ID
func GetNote(c *gin.Context) {
	userID := c.GetString("user_id")
	noteID := c.Param("id")

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	// If checklist, get items
	var checklistItems []models.ChecklistItem
	if note.NoteType == "checklist" {
		config.DB.Where("note_id = ?", noteID).Order("position ASC").Find(&checklistItems)
	}

	c.JSON(http.StatusOK, gin.H{
		"note":            note,
		"checklist_items": checklistItems,
	})
}

// UpdateNote - Update note
func UpdateNote(c *gin.Context) {
	userID := c.GetString("user_id")
	noteID := c.Param("id")

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var req models.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Title = req.Title
	note.Content = req.Content

	config.DB.Save(&note)

	c.JSON(http.StatusOK, gin.H{
		"message": "Note updated successfully",
		"note":    note,
	})
}

// DeleteNote - Delete note
func DeleteNote(c *gin.Context) {
	userID := c.GetString("user_id")
	noteID := c.Param("id")

	// Delete checklist items first
	config.DB.Where("note_id = ?", noteID).Delete(&models.ChecklistItem{})

	// Delete note
	result := config.DB.Where("id = ? AND user_id = ?", noteID, userID).
		Delete(&models.Note{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

// AddChecklistItem - Add item to checklist
func AddChecklistItem(c *gin.Context) {
	userID := c.GetString("user_id")
	noteID := c.Param("id")

	// Verify note exists and belongs to user
	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var req models.AddChecklistItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	noteUUID, _ := uuid.Parse(noteID)
	item := models.ChecklistItem{
		NoteID:   noteUUID,
		ItemText: req.ItemText,
		Position: 0, // Can be improved to get max position + 1
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item added successfully",
		"item":    item,
	})
}
