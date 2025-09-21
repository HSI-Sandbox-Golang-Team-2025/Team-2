package user_track

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track"
)

type UserTrack struct {
	basic_model.BasicModel
	UserId  uint        `json:"userId"`
	TrackId uint        `json:"trackId"`
	Track   track.Track `json:"track"`
}
