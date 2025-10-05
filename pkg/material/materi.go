package material

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"

type Material struct {
	basic_model.BasicModel
	TrackID   uint   `json:"track_id" gorm:"column:track_id"`
	ContentID uint   `json:"content_id" gorm:"column:content_id"`
	Title     string `json:"title" gorm:"column:title"`
	Body      string `json:"body" gorm:"column:body"`
}

func (Material) TableName() string {
	return "materials"
}
