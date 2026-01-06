# Community Management Ecosystem

A comprehensive Personal & Community Management platform with microservices architecture.

## 🚀 Quick Start

```bash
# Clone the repository
cd /home/vishnu/socwhiz/community

# Configure environment variables
cp .env.example .env
# Edit .env with your MSG91 credentials and other settings

# Start all services
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f
```

## 📦 Architecture

**20 Containers:**
- 4 Infrastructure: PostgreSQL, MinIO, Redis, Nginx
- 15 Microservices: Auth, Profile, Expense, Reminder, Notes, Scanner, Document, Calendar, Community, Notice, Poll, Meeting, Chat, Notification, API Gateway  
- 1 Web Admin: React frontend

**Access Points:**
- API Gateway: http://localhost:8000
- Web Admin: http://localhost:3000  
- PostgreSQL: localhost:5432
- MinIO Console: http://localhost:9001

## 🔧 Services Overview

### Core Services
| Service | Port | Status | Description |
|---------|------|--------|-------------|
| **API Gateway** | 8000 | ✅ Ready | Single entry point, routing, auth |
| **Auth Service** | 8001 | ✅ Complete | OTP & JWT with MSG91 |
| **Profile Service** | 8002 | 🚧 In Progress | User management |
| **Notification** | 8014 | 🚧 Planned | Push notifications |

### Personal Finance Services
| Service | Port | Status | Description |
|---------|------|--------|-------------|
| **Expense** | 8003 | 🚧 Planned | Expense tracking |
| **Reminder** | 8004 | 🚧 Planned | Bill reminders |
| **Notes** | 8005 | 🚧 Planned | Personal notes |
| **Scanner** | 8006 | 🚧 Planned | OCR bill scanning |
| **Document** | 8007 | 🚧 Planned | Document vault |
| **Calendar** | 8008 | 🚧 Planned | Unified calendar |

### Community Services
| Service | Port | Status | Description |
|---------|------|--------|-------------|
| **Community** | 8009 | 🚧 Planned | Community CRUD |
| **Notice** | 8010 | 🚧 Planned | Announcements |
| **Poll** | 8011 | 🚧 Planned | Polls & voting |
| **Meeting** | 8012 | 🚧 Planned | Meeting minutes |
| **Chat** | 8013 | 🚧 Planned | Real-time messaging |

## 📱 Android App

Flutter app architecture documentation available in `/docs/android-app-guide.md`

## ⚙️ Configuration

### Required Environment Variables

```env
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=community_db

# MSG91 (OTP Service)
MSG91_AUTH_KEY=your-msg91-auth-key
MSG91_SENDER_ID=your-sender-id
MSG91_TEMPLATE_ID=your-template-id

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-in-production

# MinIO
MINIO_ROOT_USER=min ioadmin
MINIO_ROOT_PASSWORD=minioadmin123

# Redis
REDIS_HOST=redis
REDIS_PORT=6379
```

## 🧪 Testing

```bash
# Test Auth Service
curl -X POST http://localhost:8000/api/auth/send-otp \
  -H "Content-Type: application/json" \
  -d '{"mobile":"9876543210"}'

# Health checks
curl http://localhost:8000/health
curl http://localhost:8001/health
```

## 📚 API Documentation

Full API documentation available at `/docs/api-reference.md`

## 🛠️ Development

```bash
# Build specific service
cd services/auth-service
go build

# Run specific service locally
go run main.go

# Rebuild and restart service
docker-compose up -d --build auth-service
```

## 📂 Project Structure

```
community/
├── services/          # 15 microservices
│   ├── auth-service/
│   ├── profile-service/
│   └── ...
├── web-admin/         # React admin panel
├── shared/            # Shared code
├── database/          # SQL schemas
├── docs/              # Documentation
├── nginx/             # Nginx config
└── docker-compose.yml
```

## 🚧 Roadmap

**Phase 1 (Current):** ✅ Infrastructure + Auth
**Phase 2:** Personal finance services
**Phase 3:** Community services
**Phase 4:** Chat & notifications
**Phase 5:** Web admin panel
**Phase 6:** Android app development
**Phase 7:** Testing & deployment

## 📄 License

Proprietary - All rights reserved

## 👥 Support

For questions or issues, contact the development team.
