package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) OpenUserPractice(
	ctx context.Context,
	userPractice *user_practice.UserPractice,
) error {
	if err := r.db.WithContext(ctx).Create(&userPractice).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) StartUserPractice(
	ctx context.Context,
	userPractice *user_practice.UserPractice,
) error {
	if err := r.db.WithContext(ctx).Create(&userPractice).Error; err != nil {
		return err
	}
	return nil
}

type GetUserPracticeCondition struct {
	ID                    uint
	UserID                uint
	ContentID             uint
	Status                user_practice.UserPracticeStatus
	UserPracticeRecordIds []uint
}

func (r *repository) GetUserPractice(
	ctx context.Context,
	userPractice *user_practice.UserPractice,
	condition *GetUserPracticeCondition,
) error {
	db := r.db.
		WithContext(ctx).
		Preload("UserPracticeRecords").
		Preload("UserPracticeRecords.Question").
		Preload("UserPracticeRecords.QuestionAnswerChoice")

	if condition.ID != 0 {
		db = db.Where("id = ?", condition.ID)
	}

	if condition.UserID != 0 {
		db = db.Where("user_id = ?", condition.UserID)
	}

	if condition.ContentID != 0 {
		db = db.Where("content_id = ?", condition.ContentID)
	}

	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.UserPracticeRecordIds != nil {
		db = db.Preload("UserPracticeRecords", "id IN ?", condition.UserPracticeRecordIds)
	}

	if err := db.First(&userPractice).Error; err != nil {
		return err
	}

	return nil
}

type GetUserPracticesCondition struct {
	ID                    uint
	UserID                uint
	ContentID             uint
	Status                user_practice.UserPracticeStatus
	UserPracticeRecordIds []uint
}

func (r *repository) GetUserPractices(
	ctx context.Context,
	userPractices *[]user_practice.UserPractice,
	condition *GetUserPracticesCondition,
) error {
	db := r.db.
		WithContext(ctx).
		Preload("UserPracticeRecords").
		Preload("UserPracticeRecords.Question").
		Preload("UserPracticeRecords.QuestionAnswerChoice")

	if condition.ID != 0 {
		db = db.Where("id = ?", condition.ID)
	}

	if condition.UserID != 0 {
		db = db.Where("user_id = ?", condition.UserID)
	}

	if condition.ContentID != 0 {
		db = db.Where("content_id = ?", condition.ContentID)
	}

	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.UserPracticeRecordIds != nil {
		db = db.Preload("UserPracticeRecords", "id IN ?", condition.UserPracticeRecordIds)
	}

	if err := db.Find(&userPractices).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateUserPractice(
	ctx context.Context,
	userPractice *user_practice.UserPractice,
) error {
	if err := r.db.WithContext(ctx).Updates(&userPractice).Error; err != nil {
		return err
	}
	return nil
}
