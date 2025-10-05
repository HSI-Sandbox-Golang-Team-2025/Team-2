package service

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"

type UserTrackService interface {
	Create(userTrack *user_track.UserTrack) error
	GetByID(id uint) (*user_track.UserTrack, error)
	Update(userTrack *user_track.UserTrack) error
	Delete(id uint) error
	List() ([]user_track.UserTrack, error)
}
