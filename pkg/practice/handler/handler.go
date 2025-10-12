package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/practice/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/practices")

	group.Post("/", h.CreatePractice)
}

// CreatePractice godoc
// @Summary Create practice
// @Description Create a new practice
// @Tags Practice
// @Accept json
// @Produce json
// @Param data body content.Content true "Practice Data"
// @Success 201 {object} content.Content
// @Failure 400 {object} fiber.Map
// @Router /practices [post]
func (h *handler) CreatePractice(c *fiber.Ctx) error {
	var body content.Content

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.CreatePractice(context.Background(), body)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create practice success!",
		"data":    data,
	})
}
