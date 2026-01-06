package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Redis for caching/rate limiting
	config.InitRedis(cfg)
	defer config.CloseRedis()

	// Setup router
	router := gin.Default()

	// Global middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter())

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// API Gateway routes with proxying
	api := router.Group("/api")
	{
		// Auth Service (no auth required)
		auth := api.Group("/auth")
		{
			auth.POST("/send-otp", handlers.ProxyToAuth)
			auth.POST("/verify-otp", handlers.ProxyToAuth)
			auth.POST("/refresh", handlers.ProxyToAuth)
			auth.POST("/logout", middleware.JWTAuth(), handlers.ProxyToAuth)
		}

		// Profile Service (auth required)
		profile := api.Group("/profile")
		profile.Use(middleware.JWTAuth())
		{
			profile.GET("", handlers.ProxyToProfile)
			profile.PUT("", handlers.ProxyToProfile)
			profile.PUT("/settings", handlers.ProxyToProfile)
			profile.GET("/stats", handlers.ProxyToProfile)
		}

		// Expense Service (auth required)
		expenses := api.Group("/expenses")
		expenses.Use(middleware.JWTAuth())
		{
			expenses.POST("", handlers.ProxyToExpense)
			expenses.GET("", handlers.ProxyToExpense)
			expenses.GET("/:id", handlers.ProxyToExpense)
			expenses.PUT("/:id", handlers.ProxyToExpense)
			expenses.DELETE("/:id", handlers.ProxyToExpense)
			expenses.GET("/summary", handlers.ProxyToExpense)
			expenses.GET("/categories", handlers.ProxyToExpense)
		}

		// Reminder Service
		reminders := api.Group("/reminders")
		reminders.Use(middleware.JWTAuth())
		{
			reminders.POST("", handlers.ProxyToReminder)
			reminders.GET("", handlers.ProxyToReminder)
			reminders.PUT("/:id", handlers.ProxyToReminder)
			reminders.DELETE("/:id", handlers.ProxyToReminder)
			reminders.POST("/:id/mark-paid", handlers.ProxyToReminder)
			reminders.POST("/:id/snooze", handlers.ProxyToReminder)
		}

		// Notes Service
		notes := api.Group("/notes")
		notes.Use(middleware.JWTAuth())
		{
			notes.POST("", handlers.ProxyToNotes)
			notes.GET("", handlers.ProxyToNotes)
			notes.GET("/:id", handlers.ProxyToNotes)
			notes.PUT("/:id", handlers.ProxyToNotes)
			notes.DELETE("/:id", handlers.ProxyToNotes)
			notes.POST("/:id/checklist", handlers.ProxyToNotes)
		}

		// Scanner Service
		scanner := api.Group("/scanner")
		scanner.Use(middleware.JWTAuth())
		{
			scanner.GET("/bills", handlers.ProxyToScanner)
			scanner.POST("/scan", handlers.ProxyToScanner)
			scanner.POST("/bills/:id/confirm", handlers.ProxyToScanner)
		}

		// Document Service
		documents := api.Group("/documents")
		documents.Use(middleware.JWTAuth())
		{
			documents.GET("", handlers.ProxyToDocument)
			documents.POST("", handlers.ProxyToDocument)
			documents.GET("/:id", handlers.ProxyToDocument)
			documents.DELETE("/:id", handlers.ProxyToDocument)
			documents.GET("/categories", handlers.ProxyToDocument)
		}

		// Calendar Service
		calendar := api.Group("/calendar")
		calendar.Use(middleware.JWTAuth())
		{
			calendar.GET("", handlers.ProxyToCalendar)
			calendar.GET("/:date", handlers.ProxyToCalendar)
			calendar.POST("/events", handlers.ProxyToCalendar)
			calendar.PUT("/events/:id", handlers.ProxyToCalendar)
			calendar.DELETE("/events/:id", handlers.ProxyToCalendar)
		}

		// Community Service
		communities := api.Group("/communities")
		communities.Use(middleware.JWTAuth())
		{
			communities.POST("", handlers.ProxyToCommunity)
			communities.GET("", handlers.ProxyToCommunity)
			communities.GET("/:id", handlers.ProxyToCommunity)
			communities.PUT("/:id", handlers.ProxyToCommunity)
			communities.POST("/:id/invites", handlers.ProxyToCommunity)
			communities.GET("/:id/members", handlers.ProxyToCommunity)
			communities.DELETE("/:id/members/:userId", handlers.ProxyToCommunity)
			communities.POST("/:id/join", handlers.ProxyToCommunity)

			// Notice endpoints under community
			communities.POST("/:id/notices", handlers.ProxyToNotice)
			communities.GET("/:id/notices", handlers.ProxyToNotice)

			// Poll endpoints under community
			communities.POST("/:id/polls", handlers.ProxyToPoll)
			communities.GET("/:id/polls", handlers.ProxyToPoll)

			// Meeting endpoints under community
			communities.POST("/:id/minutes", handlers.ProxyToMeeting)
			communities.GET("/:id/minutes", handlers.ProxyToMeeting)

			// Community group chat
			communities.GET("/:id/group-chat", handlers.ProxyToChat)
		}

		// Notice Service (direct endpoints)
		notices := api.Group("/notices")
		notices.Use(middleware.JWTAuth())
		{
			notices.GET("/:id", handlers.ProxyToNotice)
			notices.PUT("/:id", handlers.ProxyToNotice)
			notices.DELETE("/:id", handlers.ProxyToNotice)
			notices.POST("/:id/broadcast", handlers.ProxyToNotice)
		}

		// Poll Service (direct endpoints)
		polls := api.Group("/polls")
		polls.Use(middleware.JWTAuth())
		{
			polls.GET("/:id", handlers.ProxyToPoll)
			polls.POST("/:id/vote", handlers.ProxyToPoll)
			polls.PUT("/:id/close", handlers.ProxyToPoll)
			polls.DELETE("/:id", handlers.ProxyToPoll)
		}

		// Meeting Service (direct endpoints)
		minutes := api.Group("/minutes")
		minutes.Use(middleware.JWTAuth())
		{
			minutes.GET("/:id", handlers.ProxyToMeeting)
			minutes.PUT("/:id", handlers.ProxyToMeeting)
			minutes.DELETE("/:id", handlers.ProxyToMeeting)
			minutes.POST("/:id/upload-pdf", handlers.ProxyToMeeting)
		}

		// Chat Service
		conversations := api.Group("/conversations")
		conversations.Use(middleware.JWTAuth())
		{
			conversations.POST("", handlers.ProxyToChat)
			conversations.GET("", handlers.ProxyToChat)
			conversations.GET("/:id", handlers.ProxyToChat)
			conversations.POST("/:id/messages", handlers.ProxyToChat)
			conversations.GET("/:id/messages", handlers.ProxyToChat)
		}

		// Notification Service
		notifications := api.Group("/notifications")
		notifications.Use(middleware.JWTAuth())
		{
			notifications.GET("", handlers.ProxyToNotification)
			notifications.PUT("/:id/read", handlers.ProxyToNotification)
			notifications.DELETE("/:id", handlers.ProxyToNotification)
			notifications.POST("/settings", handlers.ProxyToNotification)
		}
	}

	// WebSocket for chat (special handling)
	router.GET("/ws/chat", middleware.JWTAuth(), handlers.ProxyChatWebSocket)

	// Start server
	log.Printf("🚀 API Gateway starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
