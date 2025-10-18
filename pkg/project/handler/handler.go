package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/project/service"
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

// createProject godoc
// @Summary Create project content
// @Description Create a new project content
// @Tags Project
// @Accept json
// @Produce json
// @Param credentials body CreateProjectBody true "Create project content data"
// @Success 201 {object} CreateProjectSuccessResponse "Get project content success!"
// @Failure 500 {object} CreateProjectErrorResponse "Error"
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
