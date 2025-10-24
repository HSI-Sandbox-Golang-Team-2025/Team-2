package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/gofiber/fiber/v2"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
)

type handler struct {
	service service.Service
}

func NewHandler(
	app fiber.Router,
	m middleware.Middleware,
	s service.Service,
) {
	h := &handler{
		service: s,
	}

	group := app.Group("/contents")

	group.Post("/", m.JWT, h.CreateContent)
	group.Get("/", m.JWT, h.GetContents)
	group.Get("/:id", m.JWT, h.GetContent)
	group.Put("/:id", m.JWT, h.UpdateContent)
	group.Delete("/:id", m.JWT, h.DeleteContent)
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

// getContents godoc
// @Summary Get learning contents
// @Description Get learning contents (materials, practices and projects)
// @Tags Contents
// @Accept json
// @Produce json
// @Param trackId query string false "track.id" example(1)
// @Success 200 {object} GetContentsSuccessResponse "Get contents success!"
// @Router /contents [get]
func (h *handler) GetContents(c *fiber.Ctx) error {
	userAuth := user.User{}

	userAuth.ID = c.Locals("userId").(uint)
	userAuth.RoleID = c.Locals("roleId").(uint)

	data, err := h.service.GetContents(context.Background(), c.Queries(), &userAuth)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get content success!",
		"data":    data,
	})
}

// getContent godoc
// @Summary Get learning content
// @Description Get learning content (materials, practices and projects)
// @Tags Contents
// @Accept json
// @Produce json
// @Param id path int true "content.id" example(1)
// @Success 200 {object} GetContentSuccessResponse "Get content success!
// @Router /contents/{id} [get]
func (h *handler) GetContent(c *fiber.Ctx) error {
	userAuth := user.User{}
	userAuth.ID = c.Locals("userId").(uint)

	data, err := h.service.GetContent(context.Background(), c.Params("id"), &userAuth)

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
