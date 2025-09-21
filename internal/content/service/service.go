package service

import (
	"context"
	"strconv"

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

	trackId, _ := strconv.Atoi(queries["trackId"])

	condition := repository.GetContentsCondition{
		TrackID: uint(trackId),
		Type:    content.ContentType(queries["type"]),
	}

	if err := s.repo.GetContents(ctx, &contents, &condition); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &contents, nil
}

func (s *service) GetContent(ctx context.Context, paramId string) (*content.Content, error) {
	contents := content.Content{}

	userId := uint(15) // Temporary

	id, _ := strconv.Atoi(paramId)

	condition := repository.GetContentCondition{
		ID:     uint(id),
		UserID: userId,
	}

	if err := s.repo.GetContent(ctx, &contents, &condition); err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Content not found!")
	}

	return &contents, nil
}

func (s *service) CreateContent(ctx context.Context, c content.Content) (*content.Content, error) {
	c.Type = content.ContentTypeMaterial

	if err := s.repo.CreateContent(ctx, &c); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &c, nil
}

func (s *service) UpdateContent(ctx context.Context, c content.Content) error {
	return s.repo.UpdateContent(ctx, c)
}

func (s *service) DeleteContent(ctx context.Context, id int64) error {
	return s.repo.DeleteContent(ctx, id)
}
