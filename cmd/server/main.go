package main

import (
	"context"
	"log"
	"os"

	"go-cloud-api/internal/handler"
	"go-cloud-api/internal/middleware"
	"go-cloud-api/internal/repository"
	"go-cloud-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	// ---- init gin ----
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.ErrorHandler())
	// ---- init db connection ----
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer conn.Close(ctx)

	// ---- wiring (repo -> service -> handler) ----
	userRepo := repository.NewUserRepository(conn)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// ---- routes ----
	r.GET("/health", handler.Health)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUsersByID)
	r.POST("/users", userHandler.CreateUser)

	// ---- run ----
	r.Run(":8080")
}
