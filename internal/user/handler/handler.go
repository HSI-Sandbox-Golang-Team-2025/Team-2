package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

func NewHandler(app fiber.Router, s service.Service) {}
