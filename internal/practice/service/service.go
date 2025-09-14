package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice/repository"
	practiceQuestionRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_question/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	practiceRepo         repository.Repository
	contentRepo          contentRepository.Repository
	practiceQuestionRepo practiceQuestionRepository.Repository
}

func NewService(
	practiceRepo repository.Repository,
	contentRepo contentRepository.Repository,
) Service {
	return &service{
		practiceRepo: practiceRepo,
		contentRepo:  contentRepo,
	}
}

func (s *service) CreatePractice(ctx context.Context, p practice.Practice) (*practice.Practice, error) {
	content := content.Content{
		TrackID: p.TrackID,
		Type:    content.ContentTypePractice,
	}

	if err := s.contentRepo.CreateContent(ctx, &content); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	p.ContentID = content.ID

	if err := s.practiceRepo.CreatePractice(ctx, &p); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &p, nil
}

func (s *service) GetPractices(ctx context.Context, queries map[string]string) (*[]practice.Practice, error) {
	practices := []practice.Practice{}

	if err := s.practiceRepo.GetPractices(ctx, &practices, queries); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &practices, nil
}

func (s *service) GetPractice(ctx context.Context, paramId string) (*practice.Practice, error) {
	practice := practice.Practice{}

	if err := s.practiceRepo.GetPractice(ctx, &practice, paramId); err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Practice not found!")
	}

	return &practice, nil
}
