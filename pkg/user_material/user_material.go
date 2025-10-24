package user_material

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
)

type UserMaterial struct {
	basic_model.BasicModel
	UserID      uint               `json:"userId" example:"1"`
	ContentID   uint               `json:"contentId" example:"1"`
	UserTrackID uint               `json:"userTrackId" example:"1"`
	Status      UserMaterialStatus `json:"status" example:"opened"`
}

type UserMaterialStatus string

const (
	Opened UserMaterialStatus = "opened"
)

func (UserMaterial) TableName() string {
	return "user_materials"
}
