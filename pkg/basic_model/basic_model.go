package basic_model

import (
	"time"

	"gorm.io/gorm"
)

type BasicModel struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
