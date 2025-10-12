package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userTrackService service.UserTrackService
}

func NewHandler(app fiber.Router, s service.UserTrackService) {
	h := &handler{userTrackService: s}
	group := app.Group("/user-tracks")
	group.Post("/", h.CreateUserTrack)
	group.Get("/:id", h.GetUserTrackByID)
	group.Get("/", h.GetAllUserTrack)
	group.Put("/:id", h.UpdateUserTrack)
	group.Delete("/:id", h.DeleteUserTrack)
}

// CreateUserTrack godoc
// @Summary Create user track
// @Description Register a track for a user
// @Tags UserTrack
// @Accept json
// @Produce json
// @Param data body user_track.UserTrack true "UserTrack Data"
// @Success 201 {object} user_track.UserTrack
// @Failure 400 {object} fiber.Map
// @Router /user-tracks [post]
func (h *handler) CreateUserTrack(c *fiber.Ctx) error {
	var req user_track.UserTrack
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.userTrackService.Create(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(req)
}

// GetUserTrackByID godoc
// @Summary Get user track by ID
// @Description Get a user track by its ID
// @Tags UserTrack
// @Accept json
// @Produce json
// @Param id path int true "UserTrack ID"
// @Success 200 {object} user_track.UserTrack
// @Failure 404 {object} fiber.Map
// @Router /user-tracks/{id} [get]
func (h *handler) GetUserTrackByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	result, err := h.userTrackService.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// GetAllUserTrack godoc
// @Summary List user tracks
// @Description Get all user tracks
// @Tags UserTrack
// @Accept json
// @Produce json
// @Success 200 {array} user_track.UserTrack
// @Failure 500 {object} fiber.Map
// @Router /user-tracks [get]
func (h *handler) GetAllUserTrack(c *fiber.Ctx) error {
	result, err := h.userTrackService.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// UpdateUserTrack godoc
// @Summary Update user track
// @Description Update a user track by ID
// @Tags UserTrack
// @Accept json
// @Produce json
// @Param id path int true "UserTrack ID"
// @Param data body user_track.UserTrack true "UserTrack Data"
// @Success 200 {object} user_track.UserTrack
// @Failure 400 {object} fiber.Map
// @Router /user-tracks/{id} [put]
func (h *handler) UpdateUserTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req user_track.UserTrack
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.ID = uint(id)
	if err := h.userTrackService.Update(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

// DeleteUserTrack godoc
// @Summary Delete user track
// @Description Delete a user track by ID
// @Tags UserTrack
// @Accept json
// @Produce json
// @Param id path int true "UserTrack ID"
// @Success 204 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Router /user-tracks/{id} [delete]
func (h *handler) DeleteUserTrack(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.userTrackService.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
