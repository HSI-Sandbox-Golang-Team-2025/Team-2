package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
)

// Interface untuk repository UserTrack
type Repository interface {
	Create(ctx context.Context, userTrack *user_track.UserTrack) error
	GetByID(id uint) (*user_track.UserTrack, error)
	Update(userTrack *user_track.UserTrack) error
	Delete(id uint) error
	List() ([]user_track.UserTrack, error)
	GetUserTracks(
		ctx context.Context,
		userTrack *[]user_track.UserTrack,
		condition *GetUserTracksCondition,
	) error
	ValidateUserTrack(
		ctx context.Context,
		isValid *bool,
		condition *ValidateUserTrackCondition,
	) error
}
