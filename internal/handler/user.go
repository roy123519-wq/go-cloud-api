package handler

import (
	"go-cloud-api/internal/model"
	"go-cloud-api/internal/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []model.User{
	{ID: 1, Name: "Alice", Email: "alice@test.com"},
	{ID: 2, Name: "Bob", Email: "bob@test.com"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("INVALID_ID", "The provided ID is not valid"))
		return
	}
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, response.Success(u))
			return
		}
	}
	c.JSON(http.StatusNotFound, response.Fail("USER_NOT_FOUND", "User not found"))
}

func CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("INVALID_REQUEST", "invalid request body"))
		return
	}
	req.ID = len(users) + 1
	users = append(users, req)

	c.JSON(http.StatusCreated, response.Success(req))
}
