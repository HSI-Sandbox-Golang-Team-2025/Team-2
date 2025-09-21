package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	trackService service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{trackService: s}
	app.Post("/tracks", h.CreateTrack)
	app.Get("/tracks/:id", h.GetTrackByID)
	app.Get("/tracks", h.GetAllTrack)
	app.Put("/tracks/:id", h.UpdateTrack)
	app.Delete("/tracks/:id", h.DeleteTrack)
}

func (h *handler) CreateTrack(c *fiber.Ctx) error {
	var req track.Track
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := h.trackService.CreateTrack(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *handler) GetTrackByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	result, err := h.trackService.GetTrackByID(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *handler) GetAllTrack(c *fiber.Ctx) error {
	result, err := h.trackService.GetAllTrack(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *handler) UpdateTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req track.Track
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.ID = uint(id)
	if err := h.trackService.UpdateTrack(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) DeleteTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.trackService.DeleteTrack(c.Context(), int64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
