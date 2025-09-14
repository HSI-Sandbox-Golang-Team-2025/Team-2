package practice_question

import (
	practice_answer_choices "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_answer_choices"
	"gorm.io/gorm"
)

type PracticeQuestion struct {
	gorm.Model
	PracticeID    uint                                             `json:"practiceId"`
	Body          string                                           `json:"body"`
	AnswerChoices *[]practice_answer_choices.PracticeAnswerChoices `json:"answerChoices"`
}
