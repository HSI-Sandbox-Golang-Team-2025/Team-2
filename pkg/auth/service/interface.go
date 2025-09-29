package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Service interface {
	Login(ctx context.Context, user user.User) (*string, error)
	Register(ctx context.Context, user user.User) (*string, error)
}
