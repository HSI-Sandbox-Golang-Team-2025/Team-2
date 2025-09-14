package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice_question"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, pq *practice_question.PracticeQuestion) error {
	if err := r.db.WithContext(ctx).Create(&pq).Error; err != nil {
		return err
	}
	return nil
}
