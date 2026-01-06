package main

import (
	"auth-service/config"
	"auth-service/handlers"
	"auth-service/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	config.InitDB(cfg)
	defer config.CloseDB()

	// Initialize Redis
	config.InitRedis(cfg)
	defer config.CloseRedis()

	// Setup router
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Auth endpoints
	api := router.Group("/api/auth")
	{
		api.POST("/send-otp", handlers.SendOTP)
		api.POST("/verify-otp", handlers.VerifyOTP)
		api.POST("/refresh", middleware.JWTAuth(), handlers.RefreshToken)
		api.POST("/logout", middleware.JWTAuth(), handlers.Logout)
	}

	// Start server
	log.Printf("🚀 Auth Service starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
