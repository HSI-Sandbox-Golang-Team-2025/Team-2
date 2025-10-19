package user_practice_record

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question_answer_choice"
)

type UserPracticeRecord struct {
	basic_model.BasicModel
	UserPracticeID         uint                                         `json:"userPracticeId"`
	QuestionID             uint                                         `json:"questionId"`
	Question               *question.Question                           `json:"question" swaggertype:"string" extension:"x-nullable"`
	QuestionAnswerChoiceID *uint                                        `json:"questionAnswerChoiceId"`
	QuestionAnswerChoice   *question_answer_choice.QuestionAnswerChoice `json:"questionAnswerChoice" swaggertype:"string" extension:"x-nullable"`
	QuestionAnswerText     *string                                      `json:"questionAnswerText"`
	IsCorrect              *bool                                        `json:"isCorrect"`
}
