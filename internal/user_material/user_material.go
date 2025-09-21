package user_materials

import (
	"gorm.io/gorm"
)

type UserMaterial struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	ContentID uint   `json:"content_id" gorm:"column:content_id"`
	Status    string `json:"status" gorm:"column:status"`
}

func (UserMaterial) TableName() string {
	return "user_materials"
}
