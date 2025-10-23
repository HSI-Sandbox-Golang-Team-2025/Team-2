package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
)

type Repository interface {
	OpenUserPractice(
		ctx context.Context,
		userPractice *user_practice.UserPractice,
	) error
	StartUserPractice(
		ctx context.Context,
		userPractice *user_practice.UserPractice,
	) error
	UpdateUserPractice(
		ctx context.Context,
		userPractice *user_practice.UserPractice,
	) error
	GetUserPractices(
		ctx context.Context,
		userPractices *[]user_practice.UserPractice,
		condition *GetUserPracticesCondition,
	) error
	GetUserPractice(
		ctx context.Context,
		userPractice *user_practice.UserPractice,
		condition *GetUserPracticeCondition,
	) error
}
