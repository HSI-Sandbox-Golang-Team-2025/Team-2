package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Service interface {
	GetUsers(
		ctx context.Context,
		queries map[string]string,
		userAuth *user.User,
	) (*[]user.User, error)
}
