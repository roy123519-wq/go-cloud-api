package handler

import (
	"go-cloud-api/internal/repository"
	"go-cloud-api/internal/service"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := repository.NewInMemoryUserRepository(nil)
	svc := service.NewUserService(repo)
	h := NewUserHandler(svc)

	r := gin.New()
	r.GET("/users", h.GetUsers)
	r.GET("/users/:id", h.GetUsersByID)
	r.POST("/users", h.CreateUser)

	return r
}
