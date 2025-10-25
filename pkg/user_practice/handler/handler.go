package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/service"
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

	group := app.Group("/user-practices")

	group.Post("/", m.JWT, h.StartUserPractice)
	group.Get("/", m.JWT, h.GetUserPractices)
	group.Patch("/:id/submit", m.JWT, h.SubmitUserPractice)
	group.Patch("/:id/review", m.JWT, h.ReviewUserPractice)
}

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-practices [get]
func (h *handler) GetUserPractices(c *fiber.Ctx) error {
	user := user.User{}
	user.ID = c.Locals("userId").(uint)
	user.RoleID = c.Locals("roleId").(uint)

	data, err := h.service.GetUserPractices(context.Background(), c.Queries(), &user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user practices success!",
		"data":    data,
	})
}

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-practices [post]
func (h *handler) StartUserPractice(c *fiber.Ctx) error {
	body := user_practice.UserPractice{}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user := user.User{}
	user.ID = c.Locals("userId").(uint)

	data, err := h.service.StartUserPractice(context.Background(), body, user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user practice success!",
		"data":    data,
	})
}

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-practices/:id/submit [patch]
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

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-practices/{id}/review [patch]
func (h *handler) ReviewUserPractice(c *fiber.Ctx) error {
	var body user_practice.UserPractice

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user := user.User{}
	user.ID = c.Locals("userId").(uint)

	data, err := h.service.ReviewUserPractice(
		context.Background(),
		body,
		c.Params("id"),
		user,
	)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user practice success!",
		"data":    data,
	})
}
