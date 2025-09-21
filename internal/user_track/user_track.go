package user_track

import (
	"gorm.io/gorm"
)

type UserTrack struct {
	gorm.Model
	UserID  uint `json:"user_id" gorm:"column:user_id"`
	TrackID uint `json:"track_id" gorm:"column:track_id"`
}

func (UserTrack) TableName() string {
	return "user_tracks"
}
