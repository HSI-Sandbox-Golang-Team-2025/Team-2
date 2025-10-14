package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
)

type Repository interface {
	GetRole(ctx context.Context, rol *role.Role, condition *FindRoleCondition) error
	CreateRole(ctx context.Context, rol *role.Role) error
}
