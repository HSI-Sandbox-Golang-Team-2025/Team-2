package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/service"
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

	group := app.Group("/users")

	group.Get("/", m.JWT, h.GetUsers)
}

// getUserPractices godoc
// @Summary Get users
// @Description Get users
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /users [get]
func (h *handler) GetUsers(c *fiber.Ctx) error {
	userAuth := user.User{}
	userAuth.ID = c.Locals("userId").(uint)

	data, err := h.service.GetUsers(context.Background(), c.Queries(), &userAuth)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get users success!",
		"data":    data,
	})
}
