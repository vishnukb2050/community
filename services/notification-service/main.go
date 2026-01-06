package main

import (
	"log"
	"notification-service/config"
	"notification-service/handlers"
	"notification-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/notifications")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetNotifications)
		api.POST("", handlers.SendNotification)
		api.POST("/:id/read", handlers.MarkAsRead)
		api.POST("/read-all", handlers.MarkAllAsRead)
		api.DELETE("/:id", handlers.DeleteNotification)
		api.POST("/register-device", handlers.RegisterDevice)
	}

	log.Printf("🚀 notification service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
