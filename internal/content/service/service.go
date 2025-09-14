package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	repo repository.Repository
}

func NewService(practiceRepo repository.Repository) Service {
	return &service{
		repo: practiceRepo,
	}
}

func (s *service) GetContents(ctx context.Context, queries map[string]string) (*[]content.Content, error) {
	contents := []content.Content{}

	if err := s.repo.GetContents(ctx, &contents, queries); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &contents, nil
}

func (s *service) GetContent(ctx context.Context, paramId string) (*content.Content, error) {
	contents := content.Content{}

	if err := s.repo.GetContent(ctx, &contents, paramId); err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Content not found!")
	}

	return &contents, nil
}
