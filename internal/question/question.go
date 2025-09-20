package question

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question_answer_choice"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ContentID     uint                                           `json:"contentId" gorm:"unique"`
	Question      string                                         `json:"question"`
	AnswerChoices *[]question_answer_choice.QuestionAnswerChoice `json:"answerChoices"`
}
