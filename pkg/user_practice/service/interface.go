package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
)

type Service interface {
	StartUserPractice(
		ctx context.Context,
		paramId string,
	) (*user_practice.UserPractice, error)
	SubmitUserPractice(
		ctx context.Context,
		up user_practice.UserPractice,
		paramId string,
	) (*user_practice.UserPractice, error)
	ReviewUserPractice(
		ctx context.Context,
		up user_practice.UserPractice,
		paramId string,
	) (*user_practice.UserPractice, error)
	GetUserPractices(
		ctx context.Context,
		queries map[string]string,
	) (*[]user_practice.UserPractice, error)
}
