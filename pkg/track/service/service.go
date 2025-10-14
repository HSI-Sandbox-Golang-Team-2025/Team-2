package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateTrack(ctx context.Context, t track.Track) (*track.Track, error) {
	if err := s.repository.CreateTrack(ctx, &t); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return &t, nil
}

func (s *service) GetTrackByID(ctx context.Context, id int64) (*track.Track, error) {
	return s.repository.GetTrackByID(ctx, id)
}

func (s *service) GetAllTrack(ctx context.Context) ([]track.Track, error) {
	return s.repository.GetAllTrack(ctx)
}

func (s *service) UpdateTrack(ctx context.Context, t track.Track) error {
	return s.repository.UpdateTrack(ctx, t)
}

func (s *service) DeleteTrack(ctx context.Context, id int64) error {
	return s.repository.DeleteTrack(ctx, id)
}
