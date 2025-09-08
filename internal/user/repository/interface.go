package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
)

type Repository interface {
	CreateUser(ctx context.Context, u user.User) (*user.User, error)
}
