package question_answer_choice

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"

type QuestionAnswerChoice struct {
	basic_model.BasicModel
	QuestionID uint   `json:"questionId"`
	Answer     string `json:"answer"`
	IsCorrect  bool   `json:"isCorrect"`
}
