package practice_answer_choices

import (
	"gorm.io/gorm"
)

type PracticeAnswerChoices struct {
	gorm.Model
	PracticeQuestionID uint   `json:"questionId"`
	Body               string `json:"body"`
	IsCorrect          bool   `json:"isCorrect"`
}
