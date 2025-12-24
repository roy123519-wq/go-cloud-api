# go-cloud-api

A simple Go Web API project built with Gin, following clean architecture principles.

---

##  Features

- RESTful Web API using Gin
- Layered architecture (Handler / Service / Repository)
- Centralized error handling via middleware
- Unified API response format
- In-memory repository for testing
- PostgreSQL-ready repository design
- Unit tests for service & handler layers

---

##  Project Structure

cmd/server
└─ main.go # Application entry point

internal
├─ handler # HTTP layer (Gin handlers)
├─ service # Business logic layer
├─ repository # Data access layer (Postgres / InMemory)
├─ middleware # Gin middlewares (logging, error handling)
├─ model # Domain models
└─ response # Unified API response & AppError


---

##  Request Flow

### Successful request

Client
→ Gin Router
→ Middleware
→ Handler
→ Service
→ Repository
→ Service
→ Handler
→ JSON Response (200/201)


### Error request (e.g. user not found)

Repository (ErrUserNotFound)
→ Service (translate error)
→ Handler (c.Error(AppError))
→ ErrorHandler Middleware
→ JSON Response (404)


---

##  API Response Format

### Success
```json
{
  "data": {...},
  "error": null
}

### Error

{
  "data": null,
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "user not found"
  }
}

Testing

Run all tests:

go test ./...

