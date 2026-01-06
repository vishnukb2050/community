package main

import (
	"log"
	"profile-service/config"
	"profile-service/handlers"
	"profile-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())

	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/profile")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetProfile)
		api.PUT("", handlers.UpdateProfile)
		api.PUT("/settings", handlers.UpdateSettings)
		api.GET("/stats", handlers.GetStats)
	}

	log.Printf("🚀 Profile Service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
