package practice_question

import (
	"gorm.io/gorm"
)

type PracticeQuestion struct {
	gorm.Model
	PracticeID uint   `json:"practice_id"`
	Body       string `json:"body"`
}
