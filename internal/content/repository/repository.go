package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type GetContentsCondition struct {
	TrackId uint
	Type    content.ContentType
}

func (r *repository) GetContents(ctx context.Context, c *[]content.Content, condition *GetContentsCondition) error {
	err := r.db.WithContext(ctx).
		Select("*, 'Hidden' as \"body\"").
		Where(&condition).
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
		Preload("Questions").
		Preload("Questions.AnswerChoices").
		Joins("JOIN user_tracks ut ON ut.track_id = contents.track_id AND ut.user_id = ?", condition.UserID).
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
