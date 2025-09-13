package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track"
)

type Service interface {
	CreateTrack(ctx context.Context, t track.Track) (*track.Track, error)
	GetTrackByID(ctx context.Context, id int64) (*track.Track, error)
	GetAllTrack(ctx context.Context) ([]track.Track, error)
	UpdateTrack(ctx context.Context, t track.Track) error
	DeleteTrack(ctx context.Context, id int64) error
}
