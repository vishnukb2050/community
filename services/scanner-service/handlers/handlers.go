package handlers

import (
	"encoding/base64"
	"net/http"
	"regexp"
	"scanner-service/config"
	"scanner-service/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "scanner-service",
		"time":    time.Now(),
	})
}

// GetScannedBills - List all scanned bills for user
func GetScannedBills(c *gin.Context) {
	userID := c.GetString("user_id")

	var bills []models.ScannedBill
	if err := config.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&bills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bills"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bills": bills, "count": len(bills)})
}

// ScanBill - Process bill image and extract data
func ScanBill(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.ScanBillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode base64 image
	imageData, err := base64.StdEncoding.DecodeString(req.ImageBase64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}

	// Here you would normally call OCR service (Tesseract, Google Vision, etc.)
	// For now, we'll simulate OCR with pattern matching
	extractedText := simulateOCR(string(imageData))

	// Extract amount, date, vendor using regex
	amount := extractAmount(extractedText)
	date := extractDate(extractedText)
	vendor := extractVendor(extractedText)
	category := matchCategory(extractedText)

	userUUID, _ := uuid.Parse(userID)
	scannedBill := models.ScannedBill{
		UserID:         userUUID,
		ImageURL:       "", // Would save to MinIO and store URL
		ExtractedText:  extractedText,
		DetectedAmount: amount,
		DetectedDate:   date,
		DetectedVendor: vendor,
		Category:       category,
		IsProcessed:    false,
	}

	if err := config.DB.Create(&scannedBill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save scan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bill scanned successfully",
		"bill":    scannedBill,
		"suggestions": gin.H{
			"amount":   amount,
			"vendor":   vendor,
			"category": category,
		},
	})
}

// ConfirmScan - User confirms and creates expense from scan
func ConfirmScan(c *gin.Context) {
	userID := c.GetString("user_id")
	billID := c.Param("id")

	var bill models.ScannedBill
	if err := config.DB.Where("id = ? AND user_id = ?", billID, userID).
		First(&bill).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scanned bill not found"})
		return
	}

	bill.IsProcessed = true
	config.DB.Save(&bill)

	c.JSON(http.StatusOK, gin.H{
		"message": "Scan confirmed",
		"bill":    bill,
	})
}

// Helper functions for text extraction

func simulateOCR(imageData string) string {
	// In production, call actual OCR service here
	// This is a simulation for demonstration
	return "Sample Bill\nTotal: $45.50\nDate: 2026-01-04\nRestaurant XYZ"
}

func extractAmount(text string) float64 {
	// Regex to find amounts like $45.50, 45.50, etc.
	re := regexp.MustCompile(`(?:₹|Rs\.?|INR|USD|\$)\s*(\d+(?:,\d{3})*(?:\.\d{2})?)`)
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		amountStr := strings.ReplaceAll(matches[1], ",", "")
		amount, _ := strconv.ParseFloat(amountStr, 64)
		return amount
	}

	// Try without currency symbol
	re = regexp.MustCompile(`\b(\d+(?:\.\d{2})?)\b`)
	matches = re.FindStringSubmatch(text)
	if len(matches) > 1 {
		amount, _ := strconv.ParseFloat(matches[1], 64)
		return amount
	}

	return 0.0
}

func extractDate(text string) time.Time {
	// Try to find date patterns
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	match := re.FindString(text)
	if match != "" {
		date, err := time.Parse("2006-01-02", match)
		if err == nil {
			return date
		}
	}
	return time.Now()
}

func extractVendor(text string) string {
	// Simple heuristic: first line often contains vendor name
	lines := strings.Split(text, "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0])
	}
	return "Unknown"
}

func matchCategory(text string) string {
	textLower := strings.ToLower(text)

	keywords := map[string][]string{
		"Food & Dining":     {"restaurant", "cafe", "food", "burger", "pizza", "dining"},
		"Groceries":         {"grocery", "supermarket", "mart", "store"},
		"Travel & Fuel":     {"fuel", "petrol", "gas", "uber", "ola", "taxi"},
		"Health & Wellness": {"pharmacy", "medical", "hospital", "doctor", "clinic"},
		"Shopping":          {"shopping", "mall", "store", "retail"},
		"Bills & Utilities": {"electricity", "water", "internet", "mobile", "bill"},
	}

	for category, words := range keywords {
		for _, word := range words {
			if strings.Contains(textLower, word) {
				return category
			}
		}
	}

	return "Other"
}
