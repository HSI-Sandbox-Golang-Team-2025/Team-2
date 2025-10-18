package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, m middleware.Middleware, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/contents/:id/questions")

	group.Get("/", m.JWT, h.GetQuestions)
}

// getContents godoc
// @Summary Get learning contents
// @Description Get learning contents (materials, practices and projects)
// @Tags Contents
// @Accept json
// @Produce json
// @Param trackId query string false "track.id" example(1)
// @Success 200 {object} GetContentsSuccessResponse "Get contents success!"
// @Router /contents/{id}/questions [get]
func (h *handler) GetQuestions(c *fiber.Ctx) error {
	user := user.User{}
	user.ID = c.Locals("userId").(uint)

	data, err := h.service.GetQuestions(context.Background(), c.Queries(), user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get content success!",
		"data":    data,
	})
}
