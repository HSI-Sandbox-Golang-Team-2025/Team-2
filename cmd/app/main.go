package main

import (
	"log"

	_ "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/docs"
	authHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/handler"
	authRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/repository"
	authService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	userHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/handler"
	userRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/repository"
	userService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Learning Platform
// @version 1.0
// @description Learning Platform API Documentations
// @host localhost:3000
// @BasePath /api
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/learning_platform"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}

	err = db.AutoMigrate(
		&user.User{},
		&role.Role{},
	)

	if err != nil {
		log.Fatal("Auto migration failed ", err)
	}

	route := app.Group("/api")

	// Create new repos
	userRepo := userRepository.NewRepository(db)
	authRepo := authRepository.NewRepository(db)

	// Create new services
	userService := userService.NewService(userRepo)
	authService := authService.NewService(authRepo, userRepo)

	// Create new handlers
	userHandler.NewHandler(route, userService)
	authHandler.NewHandler(route, authService)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
