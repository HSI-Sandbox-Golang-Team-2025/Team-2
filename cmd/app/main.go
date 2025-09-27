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
	projectHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/project/handler"
	projectService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/project/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question_answer_choice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	userHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/handler"
	userRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/repository"
	userService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice"
	userPracticeHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice/handler"
	userPracticeRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice/repository"
	userPracticeService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice_record"
	userPracticeRecordRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice_record/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project"
	userProjectHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project/handler"
	userProjectRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project/repository"
	userProjectService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project_media"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_track"
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

	db.Exec(
		fmt.Sprintf(
			"CREATE TYPE user_practice_status AS ENUM ('%s', '%s', '%s', '%s');",
			user_practice.Opened,
			user_practice.InProgress,
			user_practice.Submitted,
			user_practice.Reviewed,
		),
	)

	err = db.AutoMigrate(
		&content.Content{},
		&question.Question{},
		&question_answer_choice.QuestionAnswerChoice{},
		&role.Role{},
		&track.Track{},
		&user.User{},
		&user_practice.UserPractice{},
		&user_practice_record.UserPracticeRecord{},
		&user_project.UserProject{},
		&user_project_media.UserProjectMedia{},
		&user_track.UserTrack{},
	)

	if err != nil {
		log.Fatal("Auto migration failed ", err)
	}

	route := app.Group("/api")

	// Create new repos
	authRepo := authRepository.NewRepository(db)
	contentRepo := contentRepository.NewRepository(db)
	userPracticeRepo := userPracticeRepository.NewRepository(db)
	userPracticeRecordRepo := userPracticeRecordRepository.NewRepository(db)
	userProjectRepo := userProjectRepository.NewRepository(db)
	userRepo := userRepository.NewRepository(db)

	// Create new services
	authSvc := authService.NewService(authRepo, userRepo)
	contentSvc := contentService.NewService(contentRepo, userPracticeRepo)
	practiceSvc := practiceService.NewService(contentRepo)
	projectSvc := projectService.NewService(contentRepo)
	userPracticeSvc := userPracticeService.NewService(userPracticeRepo, userPracticeRecordRepo)
	userProjectSvc := userProjectService.NewService(userProjectRepo)
	userSvc := userService.NewService(userRepo)

	// Create new handlers
	authHandler.NewHandler(route, authSvc)
	contentHandler.NewHandler(route, contentSvc)
	practiceHandler.NewHandler(route, practiceSvc)
	projectHandler.NewHandler(route, projectSvc)
	userPracticeHandler.NewHandler(route, userPracticeSvc)
	userProjectHandler.NewHandler(route, userProjectSvc)
	userHandler.NewHandler(route, userSvc)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
