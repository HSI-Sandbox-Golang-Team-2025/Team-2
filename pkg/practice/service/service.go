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

func (s *service) CreatePractice(
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
		TrackID:   body.TrackID,
		Title:     body.Title,
		Body:      body.Body,
		Type:      content.ContentTypePractice,
		Order:     *order,
		Questions: body.Questions,
	}

	if err := s.contentRepo.CreateContent(ctx, &content); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &content, nil
}
