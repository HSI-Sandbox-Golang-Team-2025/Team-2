package content

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	TrackID uint   `json:"track_id" gorm:"column:track_id"`
	Order   int    `json:"order" gorm:"column:order"`
	Type    string `json:"type" gorm:"column:type"`
}

func (Content) TableName() string {
	return "contents"
}
