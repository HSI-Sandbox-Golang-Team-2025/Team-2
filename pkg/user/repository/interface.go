package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Repository interface {
	CreateUser(ctx context.Context, u *user.User) error
	GetUsers(
		ctx context.Context,
		users *[]user.User,
		condition *FindUsersCondition,
	) error
	GetUser(
		ctx context.Context,
		usr *user.User,
		condition *FindUserCondition,
	) error
}
