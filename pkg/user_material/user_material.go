package user_material

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
)

type UserMaterial struct {
	basic_model.BasicModel
	UserID    uint               `json:"userId"`
	ContentID uint               `json:"contentId"`
	Status    UserMaterialStatus `json:"status"`
}

type UserMaterialStatus string

const (
	Opened UserMaterialStatus = "opened"
)

func (UserMaterial) TableName() string {
	return "user_materials"
}
