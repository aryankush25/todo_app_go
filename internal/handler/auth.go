package handler

import (
	"todo_app_go/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// Handle user registration
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Handle user login
}
