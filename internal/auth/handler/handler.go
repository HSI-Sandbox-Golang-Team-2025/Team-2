package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/auth")

	group.Post("/login", h.Login)
	group.Post("/register", h.Register)
}

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginBody true "Login credentials"
// @Success 200 {object} SuccessLoginResponse "Login success!"
// @Failure 400 {object} InvalidLoginResponse "Invalid credentials!"
// @Router /auth/login [post]
func (h *handler) Login(c *fiber.Ctx) error {
	var body user.User

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := h.service.Login(context.Background(), body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success!",
		"data": fiber.Map{
			"token": data,
		},
	})
}

// register godoc
// @Summary User register
// @Description Register a user and return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body RegisterBody true "New user data"
// @Success 200 {object} SuccessRegisterResponse "Registration success!"
// @Failure 400 {object} InvalidRegisterResponse "Registration failed!"
// @Router /auth/register [post]
func (h *handler) Register(c *fiber.Ctx) error {
	var body user.User

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := h.service.Register(context.Background(), body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registration success!",
		"data": fiber.Map{
			"token": data,
		},
	})
}
