package main

import (
	"log"
	"notes-service/config"
	"notes-service/handlers"
	"notes-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/notes")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetNotes)
		api.POST("", handlers.CreateNote)
		api.GET("/:id", handlers.GetNote)
		api.PUT("/:id", handlers.UpdateNote)
		api.DELETE("/:id", handlers.DeleteNote)
		api.POST("/:id/checklist", handlers.AddChecklistItem)
	}

	log.Printf("🚀 notes service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
