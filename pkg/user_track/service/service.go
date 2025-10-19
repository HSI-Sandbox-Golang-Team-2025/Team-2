package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
	"github.com/gofiber/fiber/v2"
)

type userTrackService struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) UserTrackService {
	return &userTrackService{repo: repo}
}

func (s *userTrackService) Create(
	ctx context.Context,
	body *user_track.UserTrack,
	user *user.User,
) (*user_track.UserTrack, error) {
	userTrack := user_track.UserTrack{
		UserID:  user.ID,
		TrackID: body.TrackID,
	}

	err := s.repo.Create(ctx, &userTrack)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userTrack, nil
}

func (s *userTrackService) GetByID(id uint) (*user_track.UserTrack, error) {
	return s.repo.GetByID(id)
}

func (s *userTrackService) Update(userTrack *user_track.UserTrack) error {
	return s.repo.Update(userTrack)
}

func (s *userTrackService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userTrackService) List() ([]user_track.UserTrack, error) {
	return s.repo.List()
}
