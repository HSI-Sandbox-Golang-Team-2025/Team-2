package middleware

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/lib"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type middleware struct {
	db *gorm.DB
}

type Middleware interface {
	JWT(c *fiber.Ctx) error
}

func NewMiddleware(db *gorm.DB) Middleware {
	return &middleware{db: db}
}

func (m *middleware) JWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	path := c.Route().Path
	method := c.Route().Method

	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	claims, err := lib.ParseJwt(authHeader)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	u := user.User{}

	// Check the user's data in DB
	err = m.db.WithContext(context.Background()).
		Where("id = ?", claims.UserID).
		First(&u).
		Error

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	e := endpoint.Endpoint{}

	// Check endpoint and path permission
	err = m.db.WithContext(context.Background()).
		Joins("JOIN acls ON acls.endpoint_id = endpoints.id AND acls.role_id = ?", u.RoleID).
		Where("endpoints.path = ? AND endpoints.method = ?", path, method).
		First(&e).
		Error

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	c.Locals("userId", u.ID)
	c.Locals("roleId", u.RoleID)

	return c.Next()
}
