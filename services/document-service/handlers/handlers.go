package handlers

import (
	"document-service/config"
	"document-service/models"
	"encoding/base64"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "document-service",
		"time":    time.Now(),
	})
}

// GetDocuments - List all user documents
func GetDocuments(c *gin.Context) {
	userID := c.GetString("user_id")
	category := c.Query("category")

	query := config.DB.Where("user_id = ?", userID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	var documents []models.Document
	if err := query.Order("created_at DESC").Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch documents"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"documents": documents, "count": len(documents)})
}

// UploadDocument - Upload new document
func UploadDocument(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.UploadDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode base64 file
	fileData, err := base64.StdEncoding.DecodeString(req.FileBase64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format"})
		return
	}

	// In production, upload to MinIO here
	// For now, simulate file storage
	fileURL := simulateFileUpload(userID, req.FileName, fileData)
	fileType := detectFileType(req.FileName)

	userUUID, _ := uuid.Parse(userID)
	document := models.Document{
		UserID:      userUUID,
		FileName:    req.FileName,
		FileURL:     fileURL,
		FileType:    fileType,
		FileSize:    int64(len(fileData)),
		Category:    req.Category,
		Description: req.Description,
	}

	if err := config.DB.Create(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save document"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Document uploaded successfully",
		"document": document,
	})
}

// GetDocument - Get single document
func GetDocument(c *gin.Context) {
	userID := c.GetString("user_id")
	docID := c.Param("id")

	var document models.Document
	if err := config.DB.Where("id = ? AND user_id = ?", docID, userID).
		First(&document).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, document)
}

// DeleteDocument - Delete document
func DeleteDocument(c *gin.Context) {
	userID := c.GetString("user_id")
	docID := c.Param("id")

	// In production, delete from MinIO here
	result := config.DB.Where("id = ? AND user_id = ?", docID, userID).
		Delete(&models.Document{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

// GetCategories - Get document categories
func GetCategories(c *gin.Context) {
	categories := []string{
		"ID Proof",
		"Bills & Receipts",
		"Contracts",
		"Insurance",
		"Tax Documents",
		"Medical Records",
		"Property Documents",
		"Other",
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// Helper functions

func simulateFileUpload(userID, fileName string, fileData []byte) string {
	// In production: upload to MinIO and return actual URL
	// MinIO URL format: http://minio:9000/documents/{userID}/{fileName}
	return fmt.Sprintf("/documents/%s/%s", userID, fileName)
}

func detectFileType(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".pdf":
		return "pdf"
	case ".jpg", ".jpeg", ".png", ".gif":
		return "image"
	case ".doc", ".docx":
		return "doc"
	case ".xls", ".xlsx":
		return "spreadsheet"
	default:
		return "other"
	}
}
