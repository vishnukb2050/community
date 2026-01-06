#!/bin/bash

# Script to generate all remaining microservices
# This creates the complete structure for each service

set -e

BASE_DIR="/home/vishnu/socwhiz/community/services"

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Generating all microservices...${NC}"

# Service definitions: "service-name:port:description"
SERVICES=(
  "expense-service:8003:Expense tracking with categories"
  "reminder-service:8004:Bill reminders and notifications"
  "notes-service:8005:Personal notes and checklists"
  "scanner-service:8006:OCR bill scanning"
  "document-service:8007:Document vault and storage"
  "calendar-service:8008:Unified calendar"
  "community-service:8009:Community management"
  "notice-service:8010:Community announcements"
  "poll-service:8011:Polls and voting"
  "meeting-service:8012:Meeting minutes"
  "chat-service:8013:Real-time messaging"
  "notification-service:8014:Push notifications"
)

# Function to create service structure
create_service() {
  local service_name=$1
  local port=$2
  local description=$3
  
  local service_dir="$BASE_DIR/$service_name"
  
  echo -e "${GREEN}Creating $service_name (Port $port)...${NC}"
  
  # Create directories
  mkdir -p "$service_dir"/{config,handlers,models,services,middleware}
  
  # Create main.go
  cat > "$service_dir/main.go" <<'EOF'
package main

import (
	"log"
	"SERVICE_NAME/config"
	"SERVICE_NAME/handlers"
	"SERVICE_NAME/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	config.InitDB(cfg)
	defer config.CloseDB()

	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api")
	api.Use(middleware.ExtractUserID())
	{
		// Add routes here
		api.GET("/", handlers.List)
		api.POST("/", handlers.Create)
		api.GET("/:id", handlers.Get)
		api.PUT("/:id", handlers.Update)
		api.DELETE("/:id", handlers.Delete)
	}

	log.Printf("🚀 SERVICE_TITLE starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
EOF

  # Replace placeholders
  sed -i "s/SERVICE_NAME/$service_name/g" "$service_dir/main.go"
  sed -i "s/SERVICE_TITLE/${service_name//-/ }/g" "$service_dir/main.go"
  
  # Create config/config.go
  cat > "$service_dir/config/config.go" <<EOF
package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var DB *gorm.DB

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "$port"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres123"),
		DBName:     getEnv("DB_NAME", "community_db"),
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
	log.Println("✅ Database connected")
}

func CloseDB() {
	if DB != nil {
		sqlDB, _ := DB.DB()
		sqlDB.Close()
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
EOF
  
  # Create handlers/handlers.go
  cat > "$service_dir/handlers/handlers.go" <<'EOF'
package handlers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "SERVICE_NAME",
		"time":    time.Now(),
	})
}

func List(c *gin.Context) {
	// Implement list logic
	c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
}

func Create(c *gin.Context) {
	// Implement create logic
	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully"})
}

func Get(c *gin.Context) {
	// Implement get logic
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func Update(c *gin.Context) {
	// Implement update logic
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func Delete(c *gin.Context) {
	// Implement delete logic
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
EOF

  sed -i "s/SERVICE_NAME/$service_name/g" "$service_dir/handlers/handlers.go"
  
  # Create middleware/middleware.go
  cat > "$service_dir/middleware/middleware.go" <<'EOF'
package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-ID")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func ExtractUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID missing"})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
EOF
  
  # Create Dockerfile
  cat > "$service_dir/Dockerfile" <<EOF
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE $port
CMD ["./main"]
EOF
  
  # Create go.mod
  cat > "$service_dir/go.mod" <<EOF
module $service_name

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.5.0
	gorm.io/driver/postgres v1.5.4
	gorm.io/gorm v1.25.5
)
EOF
  
  # Create go.sum
  touch "$service_dir/go.sum"
  
  #  Create README.md
  cat > "$service_dir/README.md" <<EOF
# $service_name

$description

## Port
$port

## Endpoints
- GET /health - Health check
- GET /api/ - List all
- POST /api/ - Create new
- GET /api/:id - Get by ID
- PUT /api/:id - Update by ID
- DELETE /api/:id - Delete by ID

## Environment Variables
- PORT=$port
- DB_HOST=postgres
- DB_PORT=5432
- DB_USER=postgres
- DB_PASSWORD=postgres123
- DB_NAME=community_db

## Build & Run
\`\`\`bash
go mod download
go run main.go
\`\`\`
EOF
  
  echo -e "${GREEN}✓ $service_name created${NC}"
}

# Generate all services
for service_def in "${SERVICES[@]}"; do
  IFS=':' read -r service_name port description <<< "$service_def"
  create_service "$service_name" "$port" "$description"
done

echo -e "${BLUE}✅ All microservices generated successfully!${NC}"
echo -e "${BLUE}📂 Location: $BASE_DIR${NC}"
