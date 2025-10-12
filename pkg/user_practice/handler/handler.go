package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/user-practices")

	group.Get("/", h.GetUserPractices)
	group.Patch("/:id/start", h.StartUserPractice)
	group.Patch("/:id/submit", h.SubmitUserPractice)
	group.Patch("/:id/review", h.ReviewUserPractice)
}

// GetUserPractices godoc
// @Summary List user practices
// @Description Get all user practices
// @Tags UserPractice
// @Accept json
// @Produce json
// @Success 200 {array} user_practice.UserPractice
// @Failure 500 {object} fiber.Map
// @Router /user-practices [get]
func (h *handler) GetUserPractices(c *fiber.Ctx) error {
	data, err := h.service.GetUserPractices(context.Background(), c.Queries())

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user practices success!",
		"data":    data,
	})
}

// StartUserPractice godoc
// @Summary Start user practice
// @Description Start a user practice by ID
// @Tags UserPractice
// @Accept json
// @Produce json
// @Param id path int true "UserPractice ID"
// @Success 200 {object} user_practice.UserPractice
// @Failure 400 {object} fiber.Map
// @Router /user-practices/{id}/start [patch]
func (h *handler) StartUserPractice(c *fiber.Ctx) error {
	data, err := h.service.StartUserPractice(context.Background(), c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user practice success!",
		"data":    data,
	})
}

// SubmitUserPractice godoc
// @Summary Submit user practice
// @Description Submit answers for a user practice
// @Tags UserPractice
// @Accept json
// @Produce json
// @Param id path int true "UserPractice ID"
// @Param data body user_practice.UserPractice true "UserPractice Data"
// @Success 200 {object} user_practice.UserPractice
// @Failure 400 {object} fiber.Map
// @Router /user-practices/{id}/submit [patch]
func (h *handler) SubmitUserPractice(c *fiber.Ctx) error {
	var body user_practice.UserPractice

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.SubmitUserPractice(context.Background(), body, c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user practice success!",
		"data":    data,
	})
}

// ReviewUserPractice godoc
// @Summary Review user practice
// @Description Review a user practice by ID
// @Tags UserPractice
// @Accept json
// @Produce json
// @Param id path int true "UserPractice ID"
// @Param data body user_practice.UserPractice true "Review Data"
// @Success 200 {object} user_practice.UserPractice
// @Failure 400 {object} fiber.Map
// @Router /user-practices/{id}/review [patch]
func (h *handler) ReviewUserPractice(c *fiber.Ctx) error {
	var body user_practice.UserPractice

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.ReviewUserPractice(context.Background(), body, c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user practice success!",
		"data":    data,
	})
}
