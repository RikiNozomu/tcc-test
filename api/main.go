package main

import (
	"os"

	"github.com/gin-gonic/gin"

	sqlite "tcc-test/api/adapters/db/sqlite"
	handler "tcc-test/api/adapters/handlers"
	models "tcc-test/api/core/models"
	service "tcc-test/api/core/services"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load environment variables
	dbName := os.Getenv("DB_NAME")

	// Setup DB
	db, err := sqlite.New(dbName)
	if err != nil {
		panic(err)
	}

	// Ping MongoDB to verify connection
	if err := db.Health(); err != nil {
		panic(err)
	}

	db.Migrate(&models.User{})

	// Initialize Gin router
	router := gin.Default()

	// Register Healthcheck & No route
	indexHandler := handler.NewIndexHandler()
	indexHandler.IndexHandler(router)

	// Setup user repository, service, and handler
	userRepo := sqlite.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userHandler.UserRoutes(router)

	// Setup authentication service and handler
	authService := service.NewAuthService(userService)
	authHandler := handler.NewAuthHandler(authService)
	authHandler.AuthRoutes(router)

	// Start Gin server on specified port (default: 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("0.0.0.0:" + port)
}
