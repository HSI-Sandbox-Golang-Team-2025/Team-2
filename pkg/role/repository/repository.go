package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type FindRoleCondition struct {
	role.Role
}

func (r *repository) GetRole(ctx context.Context, rol *role.Role, condition *FindRoleCondition) error {
	err := r.db.WithContext(ctx).
		Where("name = ?", condition.Name).
		First(&rol).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateRole(ctx context.Context, rol *role.Role) error {
	findRoleCondition := FindRoleCondition{Role: *rol}

	if err := r.GetRole(ctx, rol, &findRoleCondition); err == nil {
		return errors.New("Role is already registered")
	}

	if err := r.db.WithContext(ctx).Create(&rol).Error; err != nil {
		return err
	}

	return nil
}
