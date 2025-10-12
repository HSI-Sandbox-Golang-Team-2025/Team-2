package handler

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{service: s}
	group := app.Group("/tracks")
	group.Post("/", h.CreateTrack)
	group.Get("/", h.GetAllTrack)
	group.Get("/:id", h.GetTrackByID)
	group.Put("/:id", h.UpdateTrack)
	group.Delete("/:id", h.DeleteTrack)
}

// CreateTrack godoc
// @Summary Create track
// @Description Create a new track
// @Tags Track
// @Accept json
// @Produce json
// @Param data body track.Track true "Track Data"
// @Success 201 {object} track.Track
// @Failure 400 {object} fiber.Map
// @Router /tracks [post]
func (h *handler) CreateTrack(c *fiber.Ctx) error {
	var body track.Track
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	data, err := h.service.CreateTrack(context.Background(), body)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Create track success!", "data": data})
}

// GetAllTrack godoc
// @Summary List tracks
// @Description Get all tracks
// @Tags Track
// @Accept json
// @Produce json
// @Success 200 {array} track.Track
// @Failure 500 {object} fiber.Map
// @Router /tracks [get]
func (h *handler) GetAllTrack(c *fiber.Ctx) error {
	data, err := h.service.GetAllTrack(context.Background())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Get tracks success!", "data": data})
}

// GetTrackByID godoc
// @Summary Get track by ID
// @Description Get a track by its ID
// @Tags Track
// @Accept json
// @Produce json
// @Param id path int true "Track ID"
// @Success 200 {object} track.Track
// @Failure 404 {object} fiber.Map
// @Router /tracks/{id} [get]
func (h *handler) GetTrackByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	data, err := h.service.GetTrackByID(context.Background(), id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Get track success!", "data": data})
}

// UpdateTrack godoc
// @Summary Update track
// @Description Update a track by ID
// @Tags Track
// @Accept json
// @Produce json
// @Param id path int true "Track ID"
// @Param data body track.Track true "Track Data"
// @Success 200 {object} track.Track
// @Failure 400 {object} fiber.Map
// @Router /tracks/{id} [put]
func (h *handler) UpdateTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var req track.Track
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	req.ID = uint(id)
	if err := h.service.UpdateTrack(context.Background(), req); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

// DeleteTrack godoc
// @Summary Delete track
// @Description Delete a track by ID
// @Tags Track
// @Accept json
// @Produce json
// @Param id path int true "Track ID"
// @Success 204 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Router /tracks/{id} [delete]
func (h *handler) DeleteTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	if err := h.service.DeleteTrack(context.Background(), int64(id)); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
