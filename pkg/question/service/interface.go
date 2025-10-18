package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Service interface {
	GetQuestions(
		ctx context.Context,
		queries map[string]string,
		user user.User,
	) (*[]question.Question, error)
}
