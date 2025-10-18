package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
)

type Repository interface {
	OpenUserPractice(
		ctx context.Context,
		up *user_practice.UserPractice,
	) error
	StartUserPractice(ctx context.Context, up *user_practice.UserPractice) error
	UpdateUserPractice(ctx context.Context, up *user_practice.UserPractice) error
	GetUserPractices(
		ctx context.Context,
		userPractices *[]user_practice.UserPractice,
		queries *GetUserPracticeCondition,
	) error
	GetUserPractice(
		ctx context.Context,
		up *user_practice.UserPractice,
		condition *GetUserPracticeCondition,
	) error
}
