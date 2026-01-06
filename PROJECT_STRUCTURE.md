# Community Management Ecosystem - Complete Project Structure

## 📁 Root Directory Structure

```
community/
├── README.md                          # Main project documentation
├── docker-compose.yml                 # All 20 containers configuration
├── .env.example                       # Environment variables template
├── .gitignore                         # Git ignore rules
│
├── database/                          # Database schemas
│   ├── init.sql                       # PostgreSQL initialization
│   └── migrations/                    # Future DB migrations
│
├── nginx/                             # Nginx reverse proxy
│   ├── nginx.conf                     # Nginx configuration
│   └── ssl/                           # SSL certificates
│
├── docs/                              # Documentation
│   ├── api-reference.md               # Complete API docs
│   ├── deployment-guide.md            # Production deployment
│   ├── development-guide.md           # Developer setup
│   └── android-app-guide.md           # Android app architecture
│
├── shared/                            # Shared code across services
│   ├── middleware/                    # Common middleware
│   │   ├── auth.go                    # JWT validation
│   │   ├── cors.go                    # CORS handler
│   │   └── logger.go                  # Request logger
│   └── models/                        # Shared data models
│       ├── user.go                    # User model
│       ├── community.go               # Community model
│       └── response.go                # Standard responses
│
├── services/                          # 15 Microservices
│   │
│   ├── api-gateway/                   # Port 8000 ✅
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── README.md
│   │
│   ├── auth-service/                  # Port 8001 ✅
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── services/
│   │   ├── models/
│   │   ├── middleware/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── README.md
│   │
│   ├── profile-service/               # Port 8002 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── expense-service/               # Port 8003 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── services/                  # Category matching logic
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── reminder-service/              # Port 8004 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── services/                  # Recurring logic
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── notes-service/                 # Port 8005 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── scanner-service/               # Port 8006 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── services/
│   │   │   ├── ocr.go                 # Text extraction
│   │   │   ├── parser.go              # Amount/date parsing
│   │   │   └── category_matcher.go   # Keyword matching
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── document-service/              # Port 8007 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── services/                  # MinIO integration
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── calendar-service/              # Port 8008 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── community-service/             # Port 8009 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── services/                  # Invitation logic
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── notice-service/                # Port 8010 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── poll-service/                  # Port 8011 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── meeting-service/               # Port 8012 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── services/                  # PDF handling
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── chat-service/                  # Port 8013 🚧
│   │   ├── main.go
│   │   ├── config/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── websocket/                 # WS handlers
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   └── notification-service/          # Port 8014 🚧
│       ├── main.go
│       ├── config/
│       ├── handlers/
│       ├── models/
│       ├── services/
│       │   ├── push.go                # Firebase push
│       │   ├── sms.go                 # MSG91 SMS
│       │   └── email.go               # Email notifications
│       ├── Dockerfile
│       └── go.mod
│
├── web-admin/                         # React Admin Panel
│   ├── public/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Dashboard/
│   │   │   ├── Community/
│   │   │   ├── Members/
│   │   │   ├── Notices/
│   │   │   ├── Polls/
│   │   │   └── Meetings/
│   │   ├── pages/
│   │   │   ├── Login.jsx
│   │   │   ├── Dashboard.jsx
│   │   │   ├── Communities.jsx
│   │   │   ├── Members.jsx
│   │   │   └── Settings.jsx
│   │   ├── services/
│   │   │   └── api.js                 # API client
│   │   ├── App.jsx
│   │   └── index.jsx
│   ├── package.json
│   ├── Dockerfile
│   └── README.md
│
└── android-app/                       # Flutter Android App
    ├── lib/
    │   ├── main.dart
    │   ├── core/
    │   │   ├── api/
    │   │   │   ├── auth_api.dart
    │   │   │   ├── expense_api.dart
    │   │   │   ├── reminder_api.dart
    │   │   │   ├── community_api.dart
    │   │   │   └── chat_api.dart
    │   │   ├── models/
    │   │   │   ├── user.dart
    │   │   │   ├── expense.dart
    │   │   │   ├── reminder.dart
    │   │   │   ├── community.dart
    │   │   │   └── message.dart
    │   │   └── utils/
    │   │       ├── ocr_helper.dart
    │   │       ├── category_matcher.dart
    │   │       └── date_parser.dart
    │   ├── features/
    │   │   ├── auth/
    │   │   │   ├── otp_screen.dart
    │   │   │   └── verify_screen.dart
    │   │   ├── personal/
    │   │   │   ├── dashboard/
    │   │   │   │   └── dashboard_screen.dart
    │   │   │   ├── expenses/
    │   │   │   │   ├── expense_list_screen.dart
    │   │   │   │   ├── add_expense_screen.dart
    │   │   │   │   └── expense_summary_screen.dart
    │   │   │   ├── reminders/
    │   │   │   │   ├── reminder_list_screen.dart
    │   │   │   │   └── add_reminder_screen.dart
    │   │   │   ├── scanner/
    │   │   │   │   ├── scanner_screen.dart
    │   │   │   │   └── verify_scan_screen.dart
    │   │   │   ├── documents/
    │   │   │   │   └── document_vault_screen.dart
    │   │   │   ├── notes/
    │   │   │   │   └── notes_screen.dart
    │   │   │   └── calendar/
    │   │   │       └── calendar_screen.dart
    │   │   ├── community/
    │   │   │   ├── hub/
    │   │   │   │   └── community_hub_screen.dart
    │   │   │   ├── noticeboard/
    │   │   │   │   └── noticeboard_screen.dart
    │   │   │   ├── polls/
    │   │   │   │   ├── polls_list_screen.dart
    │   │   │   │   └── poll_vote_screen.dart
    │   │   │   ├── minutes/
    │   │   │   │   └── minutes_list_screen.dart
    │   │   │   └── directory/
    │   │   │       └── directory_screen.dart
    │   │   └── chat/
    │   │       ├── conversations_list_screen.dart
    │   │       ├── chat_screen.dart
    │   │       └── group_chat_screen.dart
    │   └── widgets/
    │       ├── expense_card.dart
    │       ├── reminder_card.dart
    │       ├── notice_card.dart
    │       └── chat_bubble.dart
    ├── android/
    ├── pubspec.yaml
    └── README.md
```

## 🔧 Service Dependencies

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client Layer                             │
├─────────────────────────────────────────────────────────────────┤
│  Android App (Flutter)          Web Admin (React)                │
└────────────────────┬────────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                    API Gateway :8000                             │
│  • Routing  • Auth Middleware  • Rate Limiting                   │
└─┬───────┬───────┬───────┬───────┬───────┬───────┬──────┬────────┘
  │       │       │       │       │       │       │      │
  ▼       ▼       ▼       ▼       ▼       ▼       ▼      ▼
┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐
│Auth│ │Prof│ │Exp │ │Rem │ │Note│ │Scan│ │Doc │ │Cal │
│8001│ │8002│ │8003│ │8004│ │8005│ │8006│ │8007│ │8008│
└────┘ └────┘ └────┘ └────┘ └────┘ └────┘ └────┘ └────┘

┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐
│Comm│ │Noti│ │Poll│ │Meet│ │Chat│ │Ntfy│
│8009│ │8010│ │8011│ │8012│ │8013│ │8014│
└────┘ └────┘ └────┘ └────┘ └────┘ └────┘
  │       │       │       │       │       │
  └───────┴───────┴───────┴───────┴───────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Infrastructure Layer                          │
├─────────────┬──────────────────┬─────────────┬─────────────────┤
│ PostgreSQL  │      MinIO       │    Redis    │     Nginx        │
│   :5432     │   :9000/:9001    │    :6379    │    :80/:443      │
└─────────────┴──────────────────┴─────────────┴─────────────────┘
```

## 📊 Database Schema Overview

**20+ Tables:**
- `users` - User profiles
- `otp_codes` - OTP verification
- `communities` - Community organizations
- `community_members` - Membership
- `expenses` - Personal expenses
- `expense_categories` - Predefined categories
- `reminders` - Bill reminders
- `notes` - Personal notes
- `checklist_items` - Checklist items
- `documents` - File vault
- `calendar_events` - Calendar
- `notices` - Announcements
- `polls`, `poll_options`, `poll_votes` - Voting
- `meeting_minutes` - Meetings
- `conversations`, `conversation_participants`, `messages` - Chat
- `notifications` - Push notifications

## 🚀 Deployment Flow

1. **Infrastructure**
   ```bash
   docker-compose up -d postgres redis minio
   ```

2. **Backend Services**
   ```bash
   docker-compose up -d --build api-gateway auth-service profile-service ...
   ```

3. **Frontend**
   ```bash
   docker-compose up -d --build web-admin
   ```

4. **Nginx Proxy**
   ```bash
   docker-compose up -d nginx
   ```

## 📱 Android App Build

```bash
cd android-app
flutter pub get
flutter build apk --release
```

## 🔐 Environment Variables

See `.env.example` for complete list:
- Database credentials
- MSG91 API keys
- JWT secret
- MinIO credentials
- Firebase credentials
- Service URLs

## 📚 Documentation

All detailed documentation in `/docs/`:
- API Reference
- Deployment Guide
- Development Setup
- Android App Guide
