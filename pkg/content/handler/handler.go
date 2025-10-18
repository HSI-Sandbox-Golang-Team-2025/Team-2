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

// CreateContent godoc
// @Summary Create content
// @Description Create a new learning content
// @Tags Content
// @Accept json
// @Produce json
// @Param data body content.Content true "Content Data"
// @Success 201 {object} content.Content
// @Failure 400 {object} fiber.Map
// @Router /contents [post]
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

// GetContents godoc
// @Summary List contents
// @Description Get all learning contents
// @Tags Content
// @Accept json
// @Produce json
// @Success 200 {array} content.Content
// @Failure 500 {object} fiber.Map
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

// GetContent godoc
// @Summary Get content by ID
// @Description Get a learning content by its ID
// @Tags Content
// @Accept json
// @Produce json
// @Param id path int true "Content ID"
// @Success 200 {object} content.Content
// @Failure 404 {object} fiber.Map
// @Router /contents/{id} [get]
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

// UpdateContent godoc
// @Summary Update content
// @Description Update a learning content by ID
// @Tags Content
// @Accept json
// @Produce json
// @Param id path int true "Content ID"
// @Param data body content.Content true "Content Data"
// @Success 200 {object} content.Content
// @Failure 400 {object} fiber.Map
// @Router /contents/{id} [put]
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

// DeleteContent godoc
// @Summary Delete content
// @Description Delete a learning content by ID
// @Tags Content
// @Accept json
// @Produce json
// @Param id path int true "Content ID"
// @Success 204 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Router /contents/{id} [delete]
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
