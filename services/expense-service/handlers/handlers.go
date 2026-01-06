package handlers

import (
	"expense-service/config"
	"expense-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "expense-service",
		"time":    time.Now(),
	})
}

// GetExpenses - List all user expenses with optional filters
func GetExpenses(c *gin.Context) {
	userID := c.GetString("user_id")
	category := c.Query("category")
	month := c.Query("month") // YYYY-MM format

	query := config.DB.Where("user_id = ?", userID)

	if category != "" {
		query = query.Where("category = ?", category)
	}
	if month != "" {
		query = query.Where("DATE_TRUNC('month', date) = ?::date", month+"-01")
	}

	var expenses []models.Expense
	if err := query.Order("date DESC").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "count": len(expenses)})
}

// CreateExpense - Create new expense
func CreateExpense(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenseDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	userUUID, _ := uuid.Parse(userID)
	expense := models.Expense{
		UserID:      userUUID,
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
		Date:        expenseDate,
	}

	if err := config.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Expense created successfully",
		"expense": expense,
	})
}

// GetExpense - Get single expense by ID
func GetExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	expenseID := c.Param("id")

	var expense models.Expense
	if err := config.DB.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

// UpdateExpense - Update expense
func UpdateExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	expenseID := c.Param("id")

	var expense models.Expense
	if err := config.DB.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	var req models.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.Amount = req.Amount
	expense.Category = req.Category
	expense.Description = req.Description

	if req.Date != "" {
		expenseDate, err := time.Parse("2006-01-02", req.Date)
		if err == nil {
			expense.Date = expenseDate
		}
	}

	config.DB.Save(&expense)

	c.JSON(http.StatusOK, gin.H{
		"message": "Expense updated successfully",
		"expense": expense,
	})
}

// DeleteExpense - Delete expense
func DeleteExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	expenseID := c.Param("id")

	result := config.DB.Where("id = ? AND user_id = ?", expenseID, userID).
		Delete(&models.Expense{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}

// GetSummary - Get monthly expense summary with categories
func GetSummary(c *gin.Context) {
	userID := c.GetString("user_id")

	// Get total by category for current month
	type CategorySum struct {
		Category string  `json:"category"`
		Total    float64 `json:"total"`
	}

	var summary []CategorySum
	config.DB.Model(&models.Expense{}).
		Select("category, SUM(amount) as total").
		Where("user_id = ? AND DATE_TRUNC('month', date) = DATE_TRUNC('month', CURRENT_DATE)", userID).
		Group("category").
		Scan(&summary)

	// Get total for current month
	var monthTotal float64
	config.DB.Model(&models.Expense{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("user_id = ? AND DATE_TRUNC('month', date) = DATE_TRUNC('month', CURRENT_DATE)", userID).
		Scan(&monthTotal)

	c.JSON(http.StatusOK, gin.H{
		"monthly_total": monthTotal,
		"by_category":   summary,
		"month":         time.Now().Format("2006-01"),
	})
}

// GetCategories - Get predefined expense categories
func GetCategories(c *gin.Context) {
	categories := []string{
		"Food & Dining",
		"Groceries",
		"Travel & Fuel",
		"Household",
		"Health & Wellness",
		"Entertainment",
		"Shopping",
		"Bills & Utilities",
		"Education",
		"Other",
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}
