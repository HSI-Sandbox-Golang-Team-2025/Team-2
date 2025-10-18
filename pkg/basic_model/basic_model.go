package basic_model

import (
	"time"

	"gorm.io/gorm"
)

type BasicModel struct {
	ID        uint           `json:"id" extensions:"x-order=00"`
	CreatedAt time.Time      `json:"createdAt" example:"2025-10-18T14:42:51.9782263+07:00" extensions:"x-order=01"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2025-10-18T14:42:51.9782263+07:00" extensions:"x-order=02"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggertype:"string" example:"null" extensions:"x-order=03"`
}
