package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type FindUserCondition struct {
	user.User
}

func (r *repository) GetUser(ctx context.Context, u *user.User, condition *FindUserCondition) error {
	err := r.db.WithContext(ctx).
		Where("nip = ?", condition.Nip).
		First(&u).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateUser(ctx context.Context, u *user.User) error {
	findUserCondition := FindUserCondition{User: *u}

	if err := r.GetUser(ctx, u, &findUserCondition); err == nil {
		return errors.New("NIP is already registered")
	}

	if err := r.db.WithContext(ctx).Create(&u).Error; err != nil {
		return err
	}

	return nil
}
