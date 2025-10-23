package user_track

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
)

type UserTrack struct {
	basic_model.BasicModel
	UserID uint `json:"userId"`
	// User    *user.User   `json:"user"`
	TrackID      uint            `json:"trackId"`
	Track        *track.Track    `json:"track"`
	Status       UserTrackStatus `json:"status"`
	AverageScore float32         `json:"averageScore"`
}

type UserTrackStatus string

const (
	InProgress UserTrackStatus = "in progress"
	Completed  UserTrackStatus = "completed"
)
