package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(ctx context.Context, u user.User) (*user.User, error) {
	condition := user.User{
		Nip: u.Nip,
	}

	if err := r.db.WithContext(ctx).Where(&condition).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
