package question

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question_answer_choice"
)

type Question struct {
	basic_model.BasicModel
	ContentID     uint                                           `json:"contentId" gorm:"unique"`
	Question      string                                         `json:"question"`
	AnswerChoices *[]question_answer_choice.QuestionAnswerChoice `json:"answerChoices"`
}
