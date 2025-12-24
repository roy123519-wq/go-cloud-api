package handler

import (
	"net/http"
	"strconv"

	"go-cloud-api/internal/response"
	"go-cloud-api/internal/service"

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
	users, err := h.svc.GetAll(c.Request.Context())
	if err != nil {
		_ = c.Error(response.NewAppError(http.StatusInternalServerError, "INTERNAL_ERROR", "internal error"))
		return
	}
	c.JSON(http.StatusOK, response.Success(users))
}

func (h *UserHandler) GetUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(response.NewAppError(http.StatusBadRequest, "INVALID_ID", "invalid user id"))
		return
	}

	u, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrUserNotFound {
			_ = c.Error(response.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "user not found"))
			return
		}
		_ = c.Error(response.NewAppError(http.StatusInternalServerError, "INTERNAL_ERROR", "internal error"))
		return
	}

	c.JSON(http.StatusOK, response.Success(u))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(response.NewAppError(http.StatusBadRequest, "INVALID_REQUEST", "invalid request body"))
		return
	}
	if req.Name == "" || req.Email == "" {
		_ = c.Error(response.NewAppError(http.StatusBadRequest, "VALIDATION_ERROR", "name and email are required"))
		return
	}

	u, err := h.svc.Create(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		_ = c.Error(response.NewAppError(http.StatusInternalServerError, "INTERNAL_ERROR", "internal error"))
		return
	}

	c.JSON(http.StatusCreated, response.Success(u))
}
