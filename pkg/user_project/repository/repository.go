package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
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
	err := r.db.WithContext(ctx).
		Preload("Medias").
		Create(&userProject).
		Error

	if err != nil {
		return err
	}

	return nil
}

type GetUserProjectsCondition struct {
	ID                  uint
	UserID              uint
	ContentID           uint
	Status              user_project.UserProjectStatus
	StatusNot           user_project.UserProjectStatus
	UserProjectMediaIds []uint
	Limit               uint
}

func (r *repository) GetUserProjects(
	ctx context.Context,
	userProjects *[]user_project.UserProject,
	condition *GetUserProjectsCondition,
) error {
	db := r.db.
		WithContext(ctx).
		Preload("Medias")

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

	if condition.StatusNot != "" {
		db = db.Not("status = ?", condition.StatusNot)
	}

	if condition.Limit != 0 {
		db = db.Limit(int(condition.Limit))
	}

	if condition.UserProjectMediaIds != nil {
		db = db.Preload("Medias", "id IN ?", condition.UserProjectMediaIds)
	}

	return db.Find(&userProjects).Error
}

type GetUserProjectCondition struct {
	ID                  uint
	UserID              uint
	ContentID           uint
	Status              user_project.UserProjectStatus
	StatusNot           user_project.UserProjectStatus
	UserProjectMediaIds []uint
}

func (r *repository) GetUserProject(
	ctx context.Context,
	userProject *user_project.UserProject,
	condition *GetUserProjectCondition,
) error {
	userProjects := []user_project.UserProject{}

	getUserProjectsCondition := GetUserProjectsCondition{
		ID:                  condition.ID,
		UserID:              condition.UserID,
		Status:              condition.Status,
		StatusNot:           condition.StatusNot,
		ContentID:           condition.ContentID,
		UserProjectMediaIds: condition.UserProjectMediaIds,
		Limit:               1,
	}

	err := r.GetUserProjects(ctx, &userProjects, &getUserProjectsCondition)

	if err != nil || len(userProjects) < 1 {
		return errors.New("User project not found!")
	}

	*userProject = userProjects[0]

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
