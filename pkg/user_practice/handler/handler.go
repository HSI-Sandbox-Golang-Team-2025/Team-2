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

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get content success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
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

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get content success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
// @Router /user-practices/:id/start [patch]
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

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get content success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
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
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Get content success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
// @Router /user-practices/:id/review [patch]
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
