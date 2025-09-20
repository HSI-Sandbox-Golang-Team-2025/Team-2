package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice_record"
)

type Repository interface {
	GetUserPracticeRecords(
		ctx context.Context,
		userPracticeRecords *[]user_practice_record.UserPracticeRecord,
		condition GetUserPracticeRecordsCondition,
	) error
	UpdateUserPracticeRecords(
		ctx context.Context,
		userPracticeRecords *[]user_practice_record.UserPracticeRecord,
	) error
}
