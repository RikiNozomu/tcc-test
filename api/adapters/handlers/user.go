package handlers

import (
	"net/http"
	"tcc-test/api/core/services"
	"time"

	"github.com/gin-gonic/gin"

	models "tcc-test/api/core/models"
	middleware "tcc-test/api/middlewares"
	utils "tcc-test/api/utils"
)

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserHandler struct {
	userService *services.UserService // Business logic layer
}

// NewAuthHandler initializes a new AuthHandler with the given service.
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

// Login handles user login
func (h *UserHandler) Create(c *gin.Context) {
	var req UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&models.UserCreate{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func (h *UserHandler) GetMe(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(utils.ClaimsWithUser)
	user, err := h.userService.GetUser(userClaims.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func (h *UserHandler) UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	userGroup.Use(middleware.RateLimiter())
	{
		userGroup.POST("", h.Create)                                   // Public: create user
		userGroup.GET("/me", middleware.CheckBarrierHeader(), h.GetMe) // Public: get user by ID
	}
}
