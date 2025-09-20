package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/project/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/projects")

	group.Post("/", h.CreateProject)
}

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get content success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
// @Router /projects [post]
func (h *handler) CreateProject(c *fiber.Ctx) error {
	var body content.Content

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.CreateProject(context.Background(), body)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create project success!",
		"data":    data,
	})
}
