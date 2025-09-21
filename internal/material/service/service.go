package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material/repository"

	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	repository  repository.Repository
	contentRepo contentRepository.Repository
}

func NewService(
	repository repository.Repository,
	contentRepo contentRepository.Repository,
) Service {
	return &service{
		repository:  repository,
		contentRepo: contentRepo,
	}
}

func (s *service) CreateMaterial(ctx context.Context, c content.Content) (*content.Content, error) {
	c.Type = content.ContentTypeMaterial

	if err := s.contentRepo.CreateContent(ctx, &c); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &c, nil
}

func (s *service) GetMaterialByID(ctx context.Context, id int64) (*material.Material, error) {
	return s.repository.GetMaterialByID(ctx, id)
}

func (s *service) GetAllMaterial(ctx context.Context) ([]material.Material, error) {
	return s.repository.GetAllMaterial(ctx)
}

func (s *service) UpdateMaterial(ctx context.Context, m material.Material) error {
	return s.repository.UpdateMaterial(ctx, m)
}

func (s *service) DeleteMaterial(ctx context.Context, id int64) error {
	return s.repository.DeleteMaterial(ctx, id)
}
