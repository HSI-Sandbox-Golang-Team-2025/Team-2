package content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question"
)

type Content struct {
	basic_model.BasicModel
	TrackID     uint                 `json:"trackId"`
	Title       string               `json:"title"`
	Body        string               `json:"body"`
	Type        ContentType          `json:"type" gorm:"type:content_types"`
	Order       uint                 `json:"order"`
	Questions   *[]question.Question `json:"questions"`
	IsCompleted bool                 `json:"isCompleted" gorm:"-:migration"`
}

type ContentType string

const (
	ContentTypeMaterial ContentType = "material"
	ContentTypePractice ContentType = "practice"
	ContentTypeProject  ContentType = "project"
)
