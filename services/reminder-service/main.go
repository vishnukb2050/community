package main

import (
	"log"
	"reminder-service/config"
	"reminder-service/handlers"
	"reminder-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/reminders")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetReminders)
		api.POST("", handlers.CreateReminder)
		api.PUT("/:id", handlers.UpdateReminder)
		api.DELETE("/:id", handlers.DeleteReminder)
		api.POST("/:id/mark-paid", handlers.MarkPaid)
		api.POST("/:id/snooze", handlers.SnoozeReminder)
	}

	log.Printf("🚀 reminder service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
