package main

import (
	"community-service/config"
	"community-service/handlers"
	"community-service/middleware"
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

	api := router.Group("/api/communities")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetCommunities)
		api.POST("", handlers.CreateCommunity)
		api.POST("/join", handlers.JoinCommunity)
	}

	log.Printf("🚀 community service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
