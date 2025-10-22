package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	trackService service.Service
}

func NewHandler(app fiber.Router, m middleware.Middleware, s service.Service) {
	h := &handler{trackService: s}

	route := app.Group("/tracks")

	route.Post("/", m.JWT, h.CreateTrack)
	route.Get("/:id", m.JWT, h.GetTrackByID)
	route.Get("/", m.JWT, h.GetAllTrack)
	route.Put("/:id", m.JWT, h.UpdateTrack)
	route.Delete("/:id", m.JWT, h.DeleteTrack)
}

func (h *handler) CreateTrack(c *fiber.Ctx) error {
	var req track.Track

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	result, err := h.trackService.CreateTrack(c.Context(), req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create track success!",
		"data":    result,
	})
}

func (h *handler) GetTrackByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}

	result, err := h.trackService.GetTrackByID(c.Context(), int64(id))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get track success!",
		"data":    result,
	})
}

func (h *handler) GetAllTrack(c *fiber.Ctx) error {
	result, err := h.trackService.GetAllTrack(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get tracks success!",
		"data":    result,
	})
}

func (h *handler) UpdateTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}

	var req track.Track

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	req.ID = uint(id)

	result, err := h.trackService.UpdateTrack(c.Context(), req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update track success!",
		"data":    result,
	})
}

func (h *handler) DeleteTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}

	if err := h.trackService.DeleteTrack(c.Context(), int64(id)); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete track success!",
	})
}
