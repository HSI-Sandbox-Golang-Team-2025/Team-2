package service

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
)

type userTrackService struct {
	repo repository.UserTrackRepository
}

func NewService(repo repository.UserTrackRepository) UserTrackService {
	return &userTrackService{repo: repo}
}

func (s *userTrackService) Create(userTrack *user_track.UserTrack) error {
	return s.repo.Create(userTrack)
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
