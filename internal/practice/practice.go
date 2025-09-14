package practice

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_question"
	"gorm.io/gorm"
)

type Practice struct {
	gorm.Model
	TrackID   uint                                  `json:"trackId"`
	ContentID uint                                  `json:"contentId" gorm:"unique"`
	Title     string                                `json:"title"`
	Body      string                                `json:"body"`
	Questions *[]practice_question.PracticeQuestion `json:"questions"`
}
