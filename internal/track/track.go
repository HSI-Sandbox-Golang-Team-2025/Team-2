package track

import (
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	Name string `json:"name"`
}
