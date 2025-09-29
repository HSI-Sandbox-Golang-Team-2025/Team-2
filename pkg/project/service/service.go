package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
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

func (s *service) CreateProject(ctx context.Context, body content.Content) (*content.Content, error) {
	c := body
	c.Type = content.ContentTypeProject

	if err := s.contentRepo.CreateContent(ctx, &c); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &c, nil
}
