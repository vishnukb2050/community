package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Port                   string
	RedisHost              string
	RedisPort              string
	JWTSecret              string
	AuthServiceURL         string
	ProfileServiceURL      string
	ExpenseServiceURL      string
	ReminderServiceURL     string
	NotesServiceURL        string
	ScannerServiceURL      string
	DocumentServiceURL     string
	CalendarServiceURL     string
	CommunityServiceURL    string
	NoticeServiceURL       string
	PollServiceURL         string
	MeetingServiceURL      string
	ChatServiceURL         string
	NotificationServiceURL string
}

var (
	RDB *redis.Client
	ctx = context.Background()
)

func Load() *Config {
	return &Config{
		Port:                   getEnv("PORT", "8000"),
		RedisHost:              getEnv("REDIS_HOST", "localhost"),
		RedisPort:              getEnv("REDIS_PORT", "6379"),
		JWTSecret:              getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		AuthServiceURL:         getEnv("AUTH_SERVICE_URL", "http://localhost:8001"),
		ProfileServiceURL:      getEnv("PROFILE_SERVICE_URL", "http://localhost:8002"),
		ExpenseServiceURL:      getEnv("EXPENSE_SERVICE_URL", "http://localhost:8003"),
		ReminderServiceURL:     getEnv("REMINDER_SERVICE_URL", "http://localhost:8004"),
		NotesServiceURL:        getEnv("NOTES_SERVICE_URL", "http://localhost:8005"),
		ScannerServiceURL:      getEnv("SCANNER_SERVICE_URL", "http://localhost:8006"),
		DocumentServiceURL:     getEnv("DOCUMENT_SERVICE_URL", "http://localhost:8007"),
		CalendarServiceURL:     getEnv("CALENDAR_SERVICE_URL", "http://localhost:8008"),
		CommunityServiceURL:    getEnv("COMMUNITY_SERVICE_URL", "http://localhost:8009"),
		NoticeServiceURL:       getEnv("NOTICE_SERVICE_URL", "http://localhost:8010"),
		PollServiceURL:         getEnv("POLL_SERVICE_URL", "http://localhost:8011"),
		MeetingServiceURL:      getEnv("MEETING_SERVICE_URL", "http://localhost:8012"),
		ChatServiceURL:         getEnv("CHAT_SERVICE_URL", "http://localhost:8013"),
		NotificationServiceURL: getEnv("NOTIFICATION_SERVICE_URL", "http://localhost:8014"),
	}
}

func InitRedis(cfg *Config) {
	RDB = redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
		DB:   0,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("✅ Redis connected successfully")
}

func CloseRedis() {
	if RDB != nil {
		RDB.Close()
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
