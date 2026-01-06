#!/bin/bash

# Safe Deployment Script for Low-Resource Environments (e.g., EC2 Micro)
# This script builds services sequentially to prevent "No space left on device" errors.

echo "🧹 Cleaning up docker system to free space..."
docker system prune -af

echo "🏗️  Building services sequentially..."

# Core Services
echo "Building API Gateway..."
docker-compose build api-gateway || exit 1

echo "Building Auth Service..."
docker-compose build auth-service || exit 1

echo "Building Profile Service..."
docker-compose build profile-service || exit 1

# Business Logic Services
echo "Building Expense Service..."
docker-compose build expense-service || exit 1

echo "Building Reminder Service..."
docker-compose build reminder-service || exit 1

echo "Building Notes Service..."
docker-compose build notes-service || exit 1

echo "Building Scanner Service..."
docker-compose build scanner-service || exit 1

echo "Building Document Service..."
docker-compose build document-service || exit 1

echo "Building Calendar Service..."
docker-compose build calendar-service || exit 1

echo "Building Community Service..."
docker-compose build community-service || exit 1

echo "Building Notice Service..."
docker-compose build notice-service || exit 1

echo "Building Poll Service..."
docker-compose build poll-service || exit 1

echo "Building Meeting Service..."
docker-compose build meeting-service || exit 1

echo "Building Chat Service..."
docker-compose build chat-service || exit 1

echo "Building Notification Service..."
docker-compose build notification-service || exit 1

# Frontend
echo "Building Web Admin..."
docker-compose build web-admin || exit 1

echo "🚀 Starting all services..."
docker-compose up -d

echo "✅ Deployment Complete! All systems operational."
