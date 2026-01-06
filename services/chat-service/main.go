package main

import (
	"chat-service/config"
	"chat-service/handlers"
	"chat-service/middleware"
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

	api := router.Group("/api/conversations")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetConversations)
		api.POST("", handlers.CreateConversation)
		api.GET("/:id", handlers.GetConversation)
		api.GET("/:id/messages", handlers.GetMessages)
		api.POST("/:id/messages", handlers.SendMessage)
	}

	// WebSocket endpoint (implement separately if needed)
	// router.GET("/ws/chat", handlers.ChatWebSocket)

	log.Printf("🚀 chat service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
