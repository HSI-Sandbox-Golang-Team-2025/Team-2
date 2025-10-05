package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userMaterialService service.UserMaterialService
}

func NewHandler(app fiber.Router, s service.UserMaterialService) {
	h := &handler{userMaterialService: s}
	group := app.Group("/user-materials")
	group.Post("/", h.CreateUserMaterial)
	group.Get("/:id", h.GetUserMaterialByID)
	group.Get("/", h.GetAllUserMaterial)
	group.Put("/:id", h.UpdateUserMaterial)
	group.Delete("/:id", h.DeleteUserMaterial)
}

func (h *handler) CreateUserMaterial(c *fiber.Ctx) error {
	var req user_material.UserMaterial
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.userMaterialService.Create(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(req)
}

func (h *handler) GetUserMaterialByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	result, err := h.userMaterialService.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *handler) GetAllUserMaterial(c *fiber.Ctx) error {
	result, err := h.userMaterialService.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *handler) UpdateUserMaterial(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req user_material.UserMaterial
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.ID = uint(id)
	if err := h.userMaterialService.Update(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) DeleteUserMaterial(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.userMaterialService.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
