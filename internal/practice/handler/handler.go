package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/practices")

	// api/track/1/practices

	group.Post("/", h.CreatePractice)
	group.Get("/", h.GetPractices)
	group.Get("/:id", h.GetPractice)
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
// @Router /practices [post]
func (h *handler) CreatePractice(c *fiber.Ctx) error {
	var body practice.Practice

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.CreatePractice(context.Background(), body)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create practice success!",
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
// @Router /practices [get]
func (h *handler) GetPractices(c *fiber.Ctx) error {
	data, err := h.service.GetPractices(context.Background(), c.Queries())

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get practices success!",
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
// @Router /practices [get]
func (h *handler) GetPractice(c *fiber.Ctx) error {
	data, err := h.service.GetPractice(context.Background(), c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get practices success!",
		"data":    data,
	})
}
