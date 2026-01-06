package main

import (
	"document-service/config"
	"document-service/handlers"
	"document-service/middleware"
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

	api := router.Group("/api/documents")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetDocuments)
		api.POST("", handlers.UploadDocument)
		api.GET("/:id", handlers.GetDocument)
		api.DELETE("/:id", handlers.DeleteDocument)
		api.GET("/categories", handlers.GetCategories)
	}

	log.Printf("🚀 document service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
