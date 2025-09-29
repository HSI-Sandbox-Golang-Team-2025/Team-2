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

func (h *handler) GetAllMaterial(c *fiber.Ctx) error {
	result, err := h.materialService.GetAllMaterial(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

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
