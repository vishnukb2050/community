# Auth Service

OTP-based authentication service with MSG91 integration and JWT token management.

## Features

- ✅ Send OTP via MSG91 SMS
- ✅ Verify OTP and issue JWT tokens
- ✅ Refresh token support
- ✅ Token validation middleware
- ✅ User creation on first login
- ✅ Redis caching for sessions

## Environment Variables

```env
PORT=8001
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=community_db
REDIS_HOST=redis
REDIS_PORT=6379
JWT_SECRET=your-super-secret-jwt-key
MSG91_AUTH_KEY=your-msg91-auth-key
MSG91_SENDER_ID=your-sender-id
MSG91_TEMPLATE_ID=your-template-id
```

## API Endpoints

### POST /api/auth/send-otp
Send OTP to mobile number

**Request:**
```json
{
  "mobile": "9876543210"
}
```

**Response:**
```json
{
  "message": "OTP sent successfully",
  "expires_in_seconds": 300
}
```

### POST /api/auth/verify-otp
Verify OTP and get JWT token

**Request:**
```json
{
  "mobile": "9876543210",
  "otp": "123456"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOi...",
  "refresh_token": "eyJhbGciOi...",
  "user": {
    "id": "uuid",
    "mobile": "9876543210",
    "name": null,
    "email": null
  }
}
```

### POST /api/auth/refresh
Refresh access token

**Headers:**
```
Authorization: Bearer <token>
```

**Request:**
```json
{
  "refresh_token": "eyJhbGciOi..."
}
```

**Response:**
```json
{
  "token": "eyJhbGciOi..."
}
```

### POST /api/auth/logout
Logout user

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Logged out successfully"
}
```

## Build & Run

```bash
# Install dependencies
go mod download

# Run locally
go run main.go

# Build Docker image
docker build -t community-auth-service .

# Run container
docker run -p 8001:8001 --env-file .env community-auth-service
```

## MSG91 Integration

The service uses MSG91's Flow API for sending OTP SMS. Make sure to:
1. Create an MSG91 account
2. Create an OTP template with variable `{#var#}` for OTP
3. Add the auth key, sender ID, and template ID to environment variables

## JWT Token Structure

**Access Token (24 hours):**
```json
{
  "user_id": "uuid",
  "exp": 1234567890,
  "iat": 1234567890
}
```

**Refresh Token (30 days):**
```json
{
  "user_id": "uuid",
  "exp": 1234567890,
  "iat": 1234567890,
  "type": "refresh"
}
```
