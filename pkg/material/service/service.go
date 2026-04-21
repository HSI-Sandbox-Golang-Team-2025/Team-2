package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material/repository"

	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	repository  repository.Repository
	contentRepo contentRepository.Repository
}

func NewService(
	contentRepo contentRepository.Repository,
	repository repository.Repository,
) Service {
	return &service{
		repository:  repository,
		contentRepo: contentRepo,
	}
}

func (s *service) CreateMaterial(
	ctx context.Context,
	body content.Content,
) (*content.Content, error) {
	generateContentOrderCondition := contentRepository.GetContentNewOrderCondition{}
	generateContentOrderCondition.TrackId = body.TrackID

	order, err := s.contentRepo.GenerateContentOrder(
		ctx,
		&generateContentOrderCondition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	content := content.Content{
		TrackID:  body.TrackID,
		Title:    body.Title,
		Body:     body.Body,
		VideoURL: body.VideoURL,
		Type:     content.ContentTypeMaterial,
		Order:    *order,
	}

	if err := s.contentRepo.CreateContent(ctx, &content); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &content, nil
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
