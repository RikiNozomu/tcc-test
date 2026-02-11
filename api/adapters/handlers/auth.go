package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	services "tcc-test/api/core/services"
	middleware "tcc-test/api/middlewares"
)

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type AuthHandler struct {
	service *services.AuthService // Business logic layer
}

// NewAuthHandler initializes a new AuthHandler with the given service.
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token:     result.Token,
		ExpiredAt: result.ExpiredAt,
	})
}

func (h *AuthHandler) AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	authGroup.Use(middleware.RateLimiter())
	{
		authGroup.POST("/login", h.Login)
	}
}
