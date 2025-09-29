package question_answer_choice

import (
	"gorm.io/gorm"
)

type QuestionAnswerChoice struct {
	gorm.Model
	QuestionID uint   `json:"questionId"`
	Answer     string `json:"answer"`
	IsCorrect  bool   `json:"isCorrect"`
}
