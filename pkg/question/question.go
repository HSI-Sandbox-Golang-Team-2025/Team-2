package question

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question_answer_choice"
)

type Question struct {
	basic_model.BasicModel
	ContentID     uint                                           `json:"contentId" gorm:"unique" example:"1" extensions:"x-order=04"`
	Question      string                                         `json:"question" example:"Apakah 1 + 1 = 2" extensions:"x-order=05"`
	AnswerChoices *[]question_answer_choice.QuestionAnswerChoice `json:"answerChoices" extensions:"x-order=05"`
}
