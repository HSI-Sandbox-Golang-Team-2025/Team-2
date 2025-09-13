package content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	TrackID  uint               `json:"trackId"`
	Order    uint               `json:"order"`
	Type     ContentType        `json:"type" gorm:"type:content_types"`
	Practice *practice.Practice `json:"practice"`
}

type ContentType string

const (
	ContentTypeMaterial ContentType = "material"
	ContentTypePractice ContentType = "practice"
	ContentTypeProject  ContentType = "project"
)
