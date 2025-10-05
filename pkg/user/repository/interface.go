package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Repository interface {
	GetUser(ctx context.Context, u *user.User, condition *FindUserCondition) error
	CreateUser(ctx context.Context, u *user.User) error
}
