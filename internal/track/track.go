package track

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	Name      string               `json:"name"`
	Contents  *[]content.Content   `json:"-"`
	Practices *[]practice.Practice `json:"-"`
}
