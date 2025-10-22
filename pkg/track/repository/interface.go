package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
)

type Repository interface {
	CreateTrack(
		ctx context.Context,
		t *track.Track,
	) error
	GetTrack(
		ctx context.Context,
		t *track.Track,
		condition *FindTrackCondition,
	) error
	GetTrackByID(ctx context.Context, id uint, t *track.Track) error
	GetAllTrack(ctx context.Context) ([]track.Track, error)
	UpdateTrack(ctx context.Context, t *track.Track) error
	DeleteTrack(ctx context.Context, id uint) error
}
