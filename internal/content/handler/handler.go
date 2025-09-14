package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/contents")

	group.Get("/", h.GetContents)
	group.Get("/:id", h.GetContent)
}

// get contents godoc
// @Summary Get learning contents
// @Description Get learning contents (materials, practices and projects)
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get contents success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
// @Router /contents [get]
func (h *handler) GetContents(c *fiber.Ctx) error {
	data, err := h.service.GetContents(context.Background(), c.Queries())

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get content success!",
		"data":    data,
	})
}

// get content godoc
// @Summary Get learning contents
// @Description Get learning contents (materials, practices and projects)
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get contents success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
// @Router /contents [get]
func (h *handler) GetContent(c *fiber.Ctx) error {
	data, err := h.service.GetContent(context.Background(), c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get content success!",
		"data":    data,
	})
}
