package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {
	h := &handler{
		service: s,
	}

	group := app.Group("/user-projects")

	group.Post("/", h.StartUserProject)
	group.Get("/", h.GetUserProjects)
	group.Patch("/:id/submit", h.SubmitUserProject)
	group.Patch("/:id/review", h.ReviewUserProjects)
}

// StartUserProject godoc
// @Summary Start user project
// @Description Start a user project
// @Tags UserProject
// @Accept json
// @Produce json
// @Param data body user_project.UserProject true "UserProject Data"
// @Success 201 {object} user_project.UserProject
// @Failure 400 {object} fiber.Map
// @Router /user-projects [post]
func (h *handler) StartUserProject(c *fiber.Ctx) error {
	var body user_project.UserProject

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.StartUserProject(context.Background(), body)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user project success!",
		"data":    data,
	})
}

// GetUserProjects godoc
// @Summary List user projects
// @Description Get all user projects
// @Tags UserProject
// @Accept json
// @Produce json
// @Success 200 {array} user_project.UserProject
// @Failure 500 {object} fiber.Map
// @Router /user-projects [get]
func (h *handler) GetUserProjects(c *fiber.Ctx) error {
	data, err := h.service.GetUserProjects(context.Background(), c.Queries())

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user projects success!",
		"data":    data,
	})
}

// SubmitUserProject godoc
// @Summary Submit user project
// @Description Submit a user project by ID
// @Tags UserProject
// @Accept json
// @Produce json
// @Param id path int true "UserProject ID"
// @Param data body user_project.UserProject true "UserProject Data"
// @Success 200 {object} user_project.UserProject
// @Failure 400 {object} fiber.Map
// @Router /user-projects/{id}/submit [patch]
func (h *handler) SubmitUserProject(c *fiber.Ctx) error {
	var body user_project.UserProject

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.SubmitUserProject(context.Background(), body, c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user project success!",
		"data":    data,
	})
}

// ReviewUserProjects godoc
// @Summary Review user project
// @Description Review a user project by ID
// @Tags UserProject
// @Accept json
// @Produce json
// @Param id path int true "UserProject ID"
// @Param data body user_project.UserProject true "Review Data"
// @Success 200 {object} user_project.UserProject
// @Failure 400 {object} fiber.Map
// @Router /user-projects/{id}/review [patch]
func (h *handler) ReviewUserProjects(c *fiber.Ctx) error {
	var body user_project.UserProject

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	data, err := h.service.ReviewUserProject(context.Background(), body, c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user project success!",
		"data":    data,
	})
}
