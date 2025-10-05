package track

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
)

type Track struct {
	basic_model.BasicModel
	Name     string             `json:"name"`
	Contents *[]content.Content `json:"-"`
}
