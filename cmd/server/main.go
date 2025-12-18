package main

import (
	"go-cloud-api/internal/handler"
	"go-cloud-api/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUsersByID)
	r.POST("/users", handler.CreateUser)
	r.Run(":8080")
}
