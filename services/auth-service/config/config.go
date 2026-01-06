package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port            string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	RedisHost       string
	RedisPort       string
	JWTSecret       string
	MSG91AuthKey    string
	MSG91SenderID   string
	MSG91TemplateID string
}

var (
	DB  *gorm.DB
	RDB *redis.Client
	ctx = context.Background()
)

func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8001"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres123"),
		DBName:          getEnv("DB_NAME", "community_db"),
		RedisHost:       getEnv("REDIS_HOST", "localhost"),
		RedisPort:       getEnv("REDIS_PORT", "6379"),
		JWTSecret:       getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		MSG91AuthKey:    getEnv("MSG91_AUTH_KEY", ""),
		MSG91SenderID:   getEnv("MSG91_SENDER_ID", ""),
		MSG91TemplateID: getEnv("MSG91_TEMPLATE_ID", ""),
	}
}

func InitDB(cfg *Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("✅ Database connected successfully")
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

func CloseDB() {
	if DB != nil {
		sqlDB, _ := DB.DB()
		sqlDB.Close()
	}
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
