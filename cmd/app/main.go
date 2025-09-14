package main

import (
	"fmt"
	"log"

	_ "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/docs"
	authHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/handler"
	authRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/repository"
	authService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	contentHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/handler"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/repository"
	contentService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/service"
	practiceHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice/handler"
	practiceService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice/service"
	practice_answer_choices "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_answer_choices"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question_answer_choice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	userHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/handler"
	userRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/repository"
	userService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := err.Error()
			// message := "Internal server error" // Production

			if e, ok := err.(*fiber.Error); ok {
				if e.Code != fiber.StatusInternalServerError {
					code = e.Code
					message = e.Message
				}
			}

			return c.Status(code).JSON(fiber.Map{
				"message": message,
			})
		},
	})

	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/learning_platform"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}

	db.Exec(
		fmt.Sprintf(
			"CREATE TYPE content_types AS ENUM ('%s', '%s', '%s');",
			content.ContentTypeMaterial,
			content.ContentTypePractice,
			content.ContentTypeProject,
		),
	)

	err = db.AutoMigrate(
		&content.Content{},
		&practice_answer_choices.PracticeAnswerChoices{},
		&practice_question.PracticeQuestion{},
		&question.Question{},
		&question_answer_choice.QuestionAnswerChoice{},
		&role.Role{},
		&track.Track{},
		&user.User{},
	)

	if err != nil {
		log.Fatal("Auto migration failed ", err)
	}

	route := app.Group("/api")

	// Create new repos
	authRepo := authRepository.NewRepository(db)
	contentRepo := contentRepository.NewRepository(db)
	userRepo := userRepository.NewRepository(db)

	// Create new services
	authSvc := authService.NewService(authRepo, userRepo)
	contentSvc := contentService.NewService(contentRepo)
	practiceSvc := practiceService.NewService(contentRepo)
	userSvc := userService.NewService(userRepo)

	// Create new handlers
	authHandler.NewHandler(route, authSvc)
	contentHandler.NewHandler(route, contentSvc)
	practiceHandler.NewHandler(route, practiceSvc)
	userHandler.NewHandler(route, userSvc)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
