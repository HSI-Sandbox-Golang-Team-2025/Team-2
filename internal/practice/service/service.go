package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	contentRepo contentRepository.Repository
}

func NewService(
	contentRepo contentRepository.Repository,
) Service {
	return &service{
		contentRepo: contentRepo,
	}
}

func (s *service) CreatePractice(ctx context.Context, c content.Content) (*content.Content, error) {
	c.Type = content.ContentTypePractice

	if err := s.contentRepo.CreateContent(ctx, &c); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &c, nil
}
