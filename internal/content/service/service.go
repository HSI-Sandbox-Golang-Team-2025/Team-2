package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateContent(ctx context.Context, c content.Content) (*content.Content, error) {
	return s.repository.CreateContent(ctx, c)
}

func (s *service) GetContentByID(ctx context.Context, id int64) (*content.Content, error) {
	return s.repository.GetContentByID(ctx, id)
}

func (s *service) GetAllContent(ctx context.Context) ([]content.Content, error) {
	return s.repository.GetAllContent(ctx)
}

func (s *service) UpdateContent(ctx context.Context, c content.Content) error {
	return s.repository.UpdateContent(ctx, c)
}

func (s *service) DeleteContent(ctx context.Context, id int64) error {
	return s.repository.DeleteContent(ctx, id)
}
