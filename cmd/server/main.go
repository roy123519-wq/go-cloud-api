package main

import (
	"go-cloud-api/internal/handler"
	"go-cloud-api/internal/middleware"
	"go-cloud-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())
	userSvc := service.NewUserService()
	userHandler := handler.NewUserHandler(userSvc)
	r.GET("/health", handler.Health)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUsersByID)
	r.POST("/users", userHandler.CreateUser)
	r.Run(":8080")
}
