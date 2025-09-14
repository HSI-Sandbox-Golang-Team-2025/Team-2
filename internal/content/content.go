package content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	TrackID   uint                 `json:"trackId"`
	Title     string               `json:"title"`
	Body      string               `json:"body"`
	Type      ContentType          `json:"type" gorm:"type:content_types"`
	Order     uint                 `json:"order"`
	Questions *[]question.Question `json:"questions"`
}

type ContentType string

const (
	ContentTypeMaterial ContentType = "material"
	ContentTypePractice ContentType = "practice"
	ContentTypeProject  ContentType = "project"
)
