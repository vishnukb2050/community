package main

import (
	"log"
	"meeting-service/config"
	"meeting-service/handlers"
	"meeting-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/minutes")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetMinutes)
		api.POST("", handlers.CreateMinute)
		api.GET("/:id", handlers.GetMinute)
		api.DELETE("/:id", handlers.DeleteMinute)
	}

	log.Printf("🚀 meeting service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
