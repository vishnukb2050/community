# scanner-service

OCR bill scanning

## Port
8006

## Endpoints
- GET /health - Health check
- GET /api/ - List all
- POST /api/ - Create new
- GET /api/:id - Get by ID
- PUT /api/:id - Update by ID
- DELETE /api/:id - Delete by ID

## Environment Variables
- PORT=8006
- DB_HOST=postgres
- DB_PORT=5432
- DB_USER=postgres
- DB_PASSWORD=postgres123
- DB_NAME=community_db

## Build & Run
```bash
go mod download
go run main.go
```
