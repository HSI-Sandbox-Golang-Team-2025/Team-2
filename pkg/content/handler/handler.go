package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/service"
	"github.com/gofiber/fiber/v2"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/contents")

	group.Post("/", h.CreateContent)
	group.Get("/", h.GetContents)
	group.Get("/:id", h.GetContent)
	group.Put("/:id", h.UpdateContent)
	group.Delete("/:id", h.DeleteContent)
}

func (h *handler) CreateContent(c *fiber.Ctx) error {
	var req content.Content

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	data, err := h.service.CreateContent(c.Context(), req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create content success!",
		"data":    data,
	})
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

func (h *handler) UpdateContent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	var req content.Content

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	req.ID = uint(id)

	if err := h.service.UpdateContent(c.Context(), req); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) DeleteContent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	if err := h.service.DeleteContent(c.Context(), int64(id)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
