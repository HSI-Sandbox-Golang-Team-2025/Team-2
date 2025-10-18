package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/auth"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/auth/service"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
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

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body auth.LoginBody true "Login credentials"
// @Success 200 {object} auth.SuccessLoginResponse "Login success!"
// @Failure 400 {object} auth.InvalidLoginResponse "Invalid credentials!"
// @Router /auth/login [post]
func (h *handler) Login(c *fiber.Ctx) error {
	var body auth.LoginBody

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	u := user.User{
		Nip:      body.NIP,
		Password: body.Password,
	}

	data, err := h.service.Login(context.Background(), u)

	if err != nil {
		return err
	}

	token := ""
	if data != nil {
		token = *data
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success!",
		"data": fiber.Map{
			"token": token,
		},
	})
}

// Register godoc
// @Summary User register
// @Description Register a user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body auth.RegisterBody true "New user data"
// @Success 200 {object} auth.SuccessRegisterResponse "Registration success!"
// @Failure 400 {object} auth.InvalidRegisterResponse "[Error message]"
// @Router /auth/register [post]
func (h *handler) Register(c *fiber.Ctx) error {
	var body auth.RegisterBody

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	u := user.User{
		Nip:      body.NIP,
		Password: body.Password,
		Name:     body.Name,
	}

	data, err := h.service.Register(context.Background(), u)

	if err != nil {
		return err
	}

	token := ""
	if data != nil {
		token = *data
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registration success!",
		"data": fiber.Map{
			"token": token,
		},
	})
}
