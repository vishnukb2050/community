package main

import (
	"log"
	"poll-service/config"
	"poll-service/handlers"
	"poll-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/polls")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetPolls)
		api.POST("", handlers.CreatePoll)
		api.GET("/:id/options", handlers.GetPollOptions)
		api.POST("/:id/vote", handlers.Vote)
	}

	log.Printf("🚀 poll service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
