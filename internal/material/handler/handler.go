package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	materialService service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{materialService: s}
	app.Post("/materials", h.CreateMaterial)
	app.Get("/materials/:id", h.GetMaterialByID)
	app.Get("/materials", h.GetAllMaterial)
	app.Put("/materials/:id", h.UpdateMaterial)
	app.Delete("/materials/:id", h.DeleteMaterial)
}

func (h *handler) CreateMaterial(c *fiber.Ctx) error {
	var req material.Material
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := h.materialService.CreateMaterial(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
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
