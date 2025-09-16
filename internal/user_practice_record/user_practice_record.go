package user_practice_record

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/question_answer_choice"
)

type UserPracticeRecord struct {
	basic_model.BasicModel
	UserPracticeID         uint                                         `json:"userPracticeId"`
	QuestionID             uint                                         `json:"questionId"`
	Question               *question.Question                           `json:"question"`
	QuestionAnswerChoiceID uint                                         `json:"questionAnswerChoiceId"`
	QuestionAnswerChoice   *question_answer_choice.QuestionAnswerChoice `json:"questionAnswerChoice"`
	QuestionAnswerText     *string                                      `json:"questionAnswerText"`
	IsCorrect              *bool                                        `json:"isCorrect"`
}
