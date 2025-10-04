package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type GetContentsCondition struct {
	ID            uint
	TrackID       uint
	Type          content.ContentType
	OrderBefore   uint
	UserID        uint
	HideBody      bool
	OnlyCompleted bool
}

func (r *repository) GetContents(ctx context.Context, c *[]content.Content, condition *GetContentsCondition) error {
	db := r.db.WithContext(ctx).
		Model(&content.Content{})

	if condition.HideBody {
		db = db.Distinct("contents.*, 'Hidden' as \"body\", CASE WHEN COALESCE(user_practices.id, user_projects.id) IS NOT null THEN true ELSE false END AS is_completed")
	} else {
		db = db.Distinct("contents.*, CASE WHEN COALESCE(user_practices.id, user_projects.id) IS NOT null THEN true ELSE false END AS is_completed")
	}

	db = db.
		Joins("LEFT JOIN user_practices ON user_practices.content_id = contents.id AND user_practices.status = 'reviewed' AND user_practices.user_id = ?", condition.UserID).
		Joins("LEFT JOIN user_projects ON user_projects.content_id = contents.id AND user_projects.status = 'approved' AND user_projects.user_id = ?", condition.UserID).
		Preload("UserPractices", "user_id = ?", condition.UserID)

	if condition.Type != "" {
		db = db.Where("contents.type = ?", condition.Type)
	}

	if condition.TrackID != 0 {
		db = db.Where("contents.track_id = ?", condition.TrackID)
	}

	if condition.OnlyCompleted {
		db = db.Where("COALESCE(user_practices.id, user_projects.id) IS NOT NULL")
	}

	if condition.OrderBefore != 0 {
		db = db.Where("contents.order < ?", condition.OrderBefore)
	}

	err := db.
		Order("contents.order").
		Find(&c).
		Error

	if err != nil {
		return err
	}

	return nil
}

type GetContentCondition struct {
	ID     uint
	UserID uint
}

func (r *repository) GetContent(
	ctx context.Context,
	c *content.Content,
	condition *GetContentCondition,
) error {
	err := r.db.WithContext(ctx).
		Select("contents.*, CASE WHEN COALESCE(user_practices.id, user_projects.id) IS NOT null THEN true ELSE false END AS is_completed").
		Joins("JOIN user_tracks ut ON ut.track_id = contents.track_id AND ut.user_id = ?", condition.UserID).
		Joins("LEFT JOIN user_practices ON user_practices.content_id = contents.id AND user_practices.status = 'reviewed' AND user_practices.user_id = ?", condition.UserID).
		Joins("LEFT JOIN user_projects ON user_projects.content_id = contents.id AND user_projects.status = 'approved' AND user_projects.user_id = ?", condition.UserID).
		Preload("UserPractices", "user_id = ?", condition.UserID).
		Preload("UserPractices.UserPracticeRecords").
		// Preload("Questions").
		// Preload("Questions.AnswerChoices").
		Where("contents.id = ?", condition.ID).
		First(&c).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetContentByID(ctx context.Context, id int64) (*content.Content, error) {
	var c content.Content
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *repository) CreateContent(ctx context.Context, c *content.Content) error {
	if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateContent(ctx context.Context, c content.Content) error {
	return r.db.WithContext(ctx).Save(&c).Error
}

func (r *repository) DeleteContent(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&content.Content{}, id).Error
}
