package main

import (
	"expense-service/config"
	"expense-service/handlers"
	"expense-service/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/expenses")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetExpenses)
		api.POST("", handlers.CreateExpense)
		api.GET("/:id", handlers.GetExpense)
		api.PUT("/:id", handlers.UpdateExpense)
		api.DELETE("/:id", handlers.DeleteExpense)
		api.GET("/summary", handlers.GetSummary)
		api.GET("/categories", handlers.GetCategories)
	}

	log.Printf("🚀 expense service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
