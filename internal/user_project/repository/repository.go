package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) StartUserProject(
	ctx context.Context,
	userProject *user_project.UserProject,
) error {
	if err := r.db.WithContext(ctx).Create(&userProject).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserProjects(
	ctx context.Context,
	userProjects *[]user_project.UserProject,
	queries map[string]string,
) error {
	err := r.db.
		WithContext(ctx).
		Preload("UserPracticeRecords").
		Preload("UserPracticeRecords.Question").
		Preload("UserPracticeRecords.QuestionAnswerChoice").
		Find(&userProjects).
		Error

	if err != nil {
		return err
	}

	return nil
}

type GetUserProjectCondition struct {
	ID                  uint
	Status              user_project.UserProjectStatus
	UserProjectMediaIds []uint
}

func (r *repository) GetUserProject(
	ctx context.Context,
	userProject *user_project.UserProject,
	condition *GetUserProjectCondition,
) error {
	db := r.db.
		WithContext(ctx).
		Preload("Medias")

	if condition.ID != 0 {
		db = db.Where("id = ?", condition.ID)
	}

	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.UserProjectMediaIds != nil {
		db = db.Preload("Medias", "id IN ?", condition.UserProjectMediaIds)
	}

	if err := db.First(&userProject).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateUserProject(
	ctx context.Context,
	userProject *user_project.UserProject,
) error {
	if err := r.db.WithContext(ctx).Updates(&userProject).Error; err != nil {
		return err
	}
	return nil
}
