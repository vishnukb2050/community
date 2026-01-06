package main

import (
	"log"
	"scanner-service/config"
	"scanner-service/handlers"
	"scanner-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api/scanner")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("/bills", handlers.GetScannedBills)
		api.POST("/scan", handlers.ScanBill)
		api.POST("/bills/:id/confirm", handlers.ConfirmScan)
	}

	log.Printf("🚀 scanner service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
