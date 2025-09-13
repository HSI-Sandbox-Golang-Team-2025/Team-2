package content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/track"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	TrackID uint        `json:"track_id"`
	Track   track.Track `json:"track"`
	Order   uint        `json:"order"`
	Type    string      `json:"type"`
}
