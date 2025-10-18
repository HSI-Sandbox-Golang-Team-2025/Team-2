package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
)

type Repository interface {
	GetQuestions(ctx context.Context, question *[]question.Question, condition *GetQuestionCondition) error
}
