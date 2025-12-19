package handler

import (
	"go-cloud-api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(gin.H{
		"status": "ok",
	}))
}
