package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type GetUserPracticeRecordsCondition struct {
	IDs            []uint
	UserPracticeId uint
}

func (r *repository) GetUserPracticeRecords(
	ctx context.Context,
	userPracticeRecords *[]user_practice_record.UserPracticeRecord,
	condition GetUserPracticeRecordsCondition,
) error {
	db := r.db.WithContext(ctx)

	if condition.IDs != nil {
		db = db.Where("id IN ?", condition.IDs)
	}

	if condition.UserPracticeId != 0 {
		db = db.Where("user_practice_id = ?", condition.UserPracticeId)
	}

	err := db.Find(&userPracticeRecords).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateUserPracticeRecords(
	ctx context.Context,
	userPracticeRecords *[]user_practice_record.UserPracticeRecord,
) error {
	db := r.db.WithContext(ctx)

	for _, record := range *userPracticeRecords {
		if err := db.Updates(&record).Error; err != nil {
			return err
		}
	}

	return nil
}
