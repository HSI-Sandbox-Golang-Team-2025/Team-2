package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_content/service"
	"github.com/gofiber/fiber/v2"
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

	group := app.Group("/users/:userId/contents")

	group.Get("/", m.JWT, h.GetUserContents)
}

// getUserContents godoc
// @Summary Get user contents progress
// @Description Get user contents progres (materials, practices and projects)
// @Tags Contents
// @Accept json
// @Produce json
// @Router /users/{userId}/contents [get]
func (h *handler) GetUserContents(c *fiber.Ctx) error {
	userAuth := user.User{}
	userAuth.ID = c.Locals("userId").(uint)
	userAuth.RoleID = c.Locals("roleId").(uint)

	queries := c.Queries()

	paramUserId := c.Params("userId")

	data, err := h.service.GetUserContents(
		context.Background(),
		paramUserId,
		queries,
		&userAuth,
	)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user contents success!",
		"data":    data,
	})
}
