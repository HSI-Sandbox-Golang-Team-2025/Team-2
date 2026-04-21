package handler

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/middleware"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/service"
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

	group := app.Group("/user-projects")

	group.Post("/", m.JWT, h.StartUserProject)
	group.Get("/", m.JWT, h.GetUserProjects)
	group.Patch("/:id/submit", m.JWT, h.SubmitUserProject)
	group.Patch("/:id/review", m.JWT, h.ReviewUserProjects)
}

// startUserProject godoc
// @Summary Start user project
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-projects [post]
func (h *handler) StartUserProject(c *fiber.Ctx) error {
	var body user_project.UserProject

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user := user.User{}
	user.ID = c.Locals("userId").(uint)

	data, err := h.service.StartUserProject(context.Background(), body, user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user project success!",
		"data":    data,
	})
}

// getUserProjects godoc
// @Summary Get user projects
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-projects [get]
func (h *handler) GetUserProjects(c *fiber.Ctx) error {
	user := user.User{}
	user.ID = c.Locals("userId").(uint)
	user.RoleID = c.Locals("roleId").(uint)

	data, err := h.service.GetUserProjects(context.Background(), c.Queries(), user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get user projects success!",
		"data":    data,
	})
}

// submitUserProject godoc
// @Summary Submit user project
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-projects/:id/submit [patch]
func (h *handler) SubmitUserProject(c *fiber.Ctx) error {
	var body user_project.UserProject

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user := user.User{}
	user.ID = c.Locals("userId").(uint)

	data, err := h.service.SubmitUserProject(
		context.Background(),
		body,
		c.Params("id"),
		user,
	)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create user project success!",
		"data":    data,
	})
}

// reviewUserProjects godoc
// @Summary Review user projects
// @Description Authenticate user with static credentials and return JWT token
// @Tags Backlog
// @Accept json
// @Produce json
// @Router /user-projects/:id/review [patch]
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
