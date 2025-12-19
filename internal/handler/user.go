package handler

import (
	"go-cloud-api/internal/response"
	"go-cloud-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type createUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.svc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail("INTERNAL_ERROR", "internal error"))
		return
	}
	c.JSON(http.StatusOK, response.Success(users))
}

func (h *UserHandler) GetUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("INVALID_ID", "invalid user id"))
		return
	}
	u, err := h.svc.GetByID(id)
	if err == service.ErrUserNotFound {
		c.JSON(http.StatusNotFound, response.Fail("USER_NOT_FOUND", "user not found"))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail("INTERNAL_ERROR", "internal error"))
		return
	}
	c.JSON(http.StatusOK, response.Success(u))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("INVALID_REQUEST", "invalid request body"))
		return
	}
	if req.Name == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, response.Fail("VALIDATION_ERROR", "name and email are required"))
		return
	}
	u, err := h.svc.Create(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail("INTERNAL_ERROR", "internal error"))
		return
	}
	c.JSON(http.StatusCreated, response.Success(u))
}
