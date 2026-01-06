package main

import (
	"log"
	"notice-service/config"
	"notice-service/handlers"
	"notice-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/notices")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetNotices)
		api.POST("", handlers.CreateNotice)
		api.GET("/:id", handlers.GetNotice)
		api.DELETE("/:id", handlers.DeleteNotice)
	}

	log.Printf("🚀 notice service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
