package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice"
)

type Repository interface {
	StartUserPractice(ctx context.Context, up *user_practice.UserPractice) error
	UpdateUserPractice(ctx context.Context, up *user_practice.UserPractice) error
	GetUserPractices(
		ctx context.Context,
		userPractices *[]user_practice.UserPractice,
		queries map[string]string,
	) error
	GetUserPractice(
		ctx context.Context,
		up *user_practice.UserPractice,
		condition *GetUserPracticeCondition,
	) error
}
