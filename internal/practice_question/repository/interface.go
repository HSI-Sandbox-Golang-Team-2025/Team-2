package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_question"
)

type Repository interface {
	Create(ctx context.Context, pq *practice_question.PracticeQuestion) error
}
