package material

import (
	"gorm.io/gorm"
)

type Material struct {
	gorm.Model
	TrackID   uint   `json:"track_id" gorm:"column:track_id"`
	ContentID uint   `json:"content_id" gorm:"column:content_id"`
	Title     string `json:"title" gorm:"column:title"`
	Body      string `json:"body" gorm:"column:body"`
}

func (Material) TableName() string {
	return "materials"
}
