package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

// Login handles user login
func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement authentication logic
	// - Validate credentials against database
	// - Generate JWT token
	// - Return token

	c.JSON(http.StatusOK, LoginResponse{
		Token: "your-token-here",
		User:  req.Email,
	})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// TODO: Implement logout logic
	// - Invalidate token
	// - Clear session

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
