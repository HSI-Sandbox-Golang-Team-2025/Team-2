package question_answer_choice

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"

type QuestionAnswerChoice struct {
	basic_model.BasicModel
	QuestionID uint   `json:"questionId" example:"1" extensions:"x-order=04"`
	Answer     string `json:"answer" example:"Benar" extensions:"x-order=05"`
	// IsCorrect  bool   `json:"isCorrect"`
}
