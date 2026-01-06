package main

import (
	"calendar-service/config"
	"calendar-service/handlers"
	"calendar-service/middleware"
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

	api := router.Group("/api/calendar")
	api.Use(middleware.ExtractUserID())
	{
		api.GET("", handlers.GetCalendar)
		api.GET("/:date", handlers.GetEventsByDate)
		api.POST("/events", handlers.CreateEvent)
		api.PUT("/events/:id", handlers.UpdateEvent)
		api.DELETE("/events/:id", handlers.DeleteEvent)
	}

	log.Printf("🚀 calendar service starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
