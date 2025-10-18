package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type GetQuestionCondition struct {
	ID        uint
	ContentID uint
}

func (r *repository) GetQuestions(ctx context.Context, question *[]question.Question, condition *GetQuestionCondition) error {
	err := r.db.WithContext(ctx).
		Where(&condition).
		Preload("AnswerChoices").
		Find(&question).
		Error

	if err != nil {
		return err
	}

	return nil
}
