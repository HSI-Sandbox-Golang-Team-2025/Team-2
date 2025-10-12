package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	materialService service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{materialService: s}

	group := app.Group("/materials")

	group.Post("/", h.CreateMaterial)
	group.Get("/:id", h.GetMaterialByID)
	group.Get("/", h.GetAllMaterial)
	group.Put("/:id", h.UpdateMaterial)
	group.Delete("/:id", h.DeleteMaterial)
}

// CreateMaterial godoc
// @Summary Create material
// @Description Create a new material
// @Tags Material
// @Accept json
// @Produce json
// @Param data body content.Content true "Content Data"
// @Success 201 {object} content.Content
// @Failure 400 {object} fiber.Map
// @Router /materials [post]
func (h *handler) CreateMaterial(c *fiber.Ctx) error {
	var req content.Content

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	data, err := h.materialService.CreateMaterial(context.Background(), req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create material success!",
		"data":    data,
	})
}

// GetMaterialByID godoc
// @Summary Get material by ID
// @Description Get a material by its ID
// @Tags Material
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Success 200 {object} content.Content
// @Failure 404 {object} fiber.Map
// @Router /materials/{id} [get]
func (h *handler) GetMaterialByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	result, err := h.materialService.GetMaterialByID(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// GetAllMaterial godoc
// @Summary List materials
// @Description Get all materials
// @Tags Material
// @Accept json
// @Produce json
// @Success 200 {array} content.Content
// @Failure 500 {object} fiber.Map
// @Router /materials [get]
func (h *handler) GetAllMaterial(c *fiber.Ctx) error {
	result, err := h.materialService.GetAllMaterial(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// UpdateMaterial godoc
// @Summary Update material
// @Description Update a material by ID
// @Tags Material
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Param data body content.Content true "Material Data"
// @Success 200 {object} content.Content
// @Failure 400 {object} fiber.Map
// @Router /materials/{id} [put]
func (h *handler) UpdateMaterial(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req material.Material
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.ID = uint(id)
	if err := h.materialService.UpdateMaterial(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

// DeleteMaterial godoc
// @Summary Delete material
// @Description Delete a material by ID
// @Tags Material
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Success 204 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Router /materials/{id} [delete]
func (h *handler) DeleteMaterial(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.materialService.DeleteMaterial(c.Context(), int64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
