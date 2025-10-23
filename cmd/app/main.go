package main

import (
	"fmt"
	"log"

	_ "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/docs"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/lib"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/acl"
	authHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/auth/handler"
	authRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/auth/repository"
	authService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/auth/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	contentHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/handler"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	contentService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint"
	endpointRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint/repository"
	materialHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material/handler"
	materialRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material/repository"
	materialService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material/service"
	practiceHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/practice/handler"
	practiceService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/practice/service"
	projectHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/project/handler"
	projectService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/project/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
	questionHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question/handler"
	questionRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question/repository"
	questionService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question_answer_choice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	roleRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/seeder"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	trackHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/handler"
	trackRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/repository"
	trackService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	userHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/handler"
	userRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/repository"
	userService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"
	userMaterialHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material/handler"
	userMaterialRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material/repository"
	userMaterialService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	userPracticeHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/handler"
	userPracticeRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/repository"
	userPracticeService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record"
	userPracticeRecordRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
	userProjectHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/handler"
	userProjectRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/repository"
	userProjectService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project_media"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	userTrackHandler "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/handler"
	userTrackRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
	userTrackService "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/service"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
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

	godotenv.Load()

	dbConfig := lib.GetDBConfig()
	dsn := dbConfig.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
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
			"CREATE TYPE user_practice_status AS ENUM ('%s', '%s', '%s');",
			user_practice.InProgress,
			user_practice.Submitted,
			user_practice.Reviewed,
		),
	)

	db.Exec(
		fmt.Sprintf(
			"CREATE TYPE user_track_status AS ENUM ('%s', '%s');",
			user_track.InProgress,
			user_track.Completed,
		),
	)

	err = db.AutoMigrate(
		&endpoint.Endpoint{},
		&role.Role{},
		&acl.ACL{},
		&user.User{},
		&track.Track{},
		&user_track.UserTrack{},
		&content.Content{},
		&question.Question{},
		&question_answer_choice.QuestionAnswerChoice{},
		&user_practice.UserPractice{},
		&user_practice_record.UserPracticeRecord{},
		&user_project.UserProject{},
		&user_project_media.UserProjectMedia{},
		&user_material.UserMaterial{},
	)

	if err != nil {
		log.Fatal("Auto migration failed ", err)
	}

	app.Use(fiberLogger.New())

	middleware := middleware.NewMiddleware(db)

	route := app.Group("/api")

	// Create new repos
	authRepo := authRepository.NewRepository(db)
	contentRepo := contentRepository.NewRepository(db)
	endpointRepo := endpointRepository.NewRepository(db)
	materialRepo := materialRepository.NewRepository(db)
	questionRepo := questionRepository.NewRepository(db)
	roleRepo := roleRepository.NewRepository(db)
	trackRepo := trackRepository.NewRepository(db)
	userMaterialRepo := userMaterialRepository.NewRepository(db)
	userPracticeRepo := userPracticeRepository.NewRepository(db)
	userPracticeRecordRepo := userPracticeRecordRepository.NewRepository(db)
	userProjectRepo := userProjectRepository.NewRepository(db)
	userRepo := userRepository.NewRepository(db)
	userTrackRepo := userTrackRepository.NewRepository(db)

	// Create new services
	authSvc := authService.NewService(authRepo, userRepo)
	contentSvc := contentService.NewService(contentRepo, userMaterialRepo, userPracticeRepo, userTrackRepo)
	materialSvc := materialService.NewService(contentRepo, materialRepo)
	practiceSvc := practiceService.NewService(contentRepo)
	projectSvc := projectService.NewService(contentRepo)
	questionSvc := questionService.NewService(questionRepo)
	trackSvc := trackService.NewService(trackRepo)
	userMaterialSvc := userMaterialService.NewService(userMaterialRepo)
	userPracticeSvc := userPracticeService.NewService(userPracticeRepo, userPracticeRecordRepo)
	userProjectSvc := userProjectService.NewService(userProjectRepo, contentRepo)
	userSvc := userService.NewService(userRepo)
	userTrackSvc := userTrackService.NewService(userTrackRepo)

	// Create new handlers
	authHandler.NewHandler(route, authSvc)
	contentHandler.NewHandler(route, middleware, contentSvc)
	materialHandler.NewHandler(route, middleware, materialSvc)
	practiceHandler.NewHandler(route, middleware, practiceSvc)
	projectHandler.NewHandler(route, middleware, projectSvc)
	questionHandler.NewHandler(route, middleware, questionSvc)
	userMaterialHandler.NewHandler(route, userMaterialSvc)
	userPracticeHandler.NewHandler(route, middleware, userPracticeSvc)
	userProjectHandler.NewHandler(route, middleware, userProjectSvc)
	userHandler.NewHandler(route, userSvc)
	userTrackHandler.NewHandler(route, middleware, userTrackSvc)
	trackHandler.NewHandler(route, middleware, trackSvc)

	seeder.RunSeeders(
		userRepo,
		roleRepo,
		endpointRepo,
		trackRepo,
	)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
