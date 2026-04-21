package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userTrackService service.UserTrackService
}

func NewHandler(
	app fiber.Router,
	m middleware.Middleware,
	s service.UserTrackService,
) {
	h := &handler{
		userTrackService: s,
	}

	group := app.Group("/user-tracks")

	group.Post("/", m.JWT, h.CreateUserTrack)
	group.Get("/", m.JWT, h.GetUserTracks)
	group.Get("/:id", m.JWT, h.GetUserTrackByID)
	group.Put("/:id", m.JWT, h.UpdateUserTrack)
	group.Delete("/:id", m.JWT, h.DeleteUserTrack)
}

func (h *handler) CreateUserTrack(c *fiber.Ctx) error {
	var req user_track.UserTrack

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := user.User{}
	user.ID = c.Locals("userId").(uint)

	data, err := h.userTrackService.Create(context.Background(), &req, &user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create practice content success!",
		"data":    data,
	})
}

func (h *handler) GetUserTracks(c *fiber.Ctx) error {
	user := user.User{}
	user.ID = c.Locals("userId").(uint)
	user.RoleID = c.Locals("roleId").(uint)

	data, err := h.userTrackService.GetUserTracks(context.Background(), c.Queries(), &user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user tracks success!",
		"data":    data,
	})
}

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

func (h *handler) GetAllUserTrack(c *fiber.Ctx) error {
	result, err := h.userTrackService.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

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
