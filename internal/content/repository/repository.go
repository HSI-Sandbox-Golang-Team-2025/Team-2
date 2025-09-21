package repository

import (
	"context"
	"strconv"

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
		Preload("Questions").
		Preload("Questions.AnswerChoices").
		Where(&condition).
		Find(&c).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetContent(ctx context.Context, c *content.Content, paramId string) error {
	id, _ := strconv.Atoi(paramId)

	condition := content.Content{}
	condition.ID = uint(id)

	if err := r.db.WithContext(ctx).Preload("Questions").Preload("Questions.AnswerChoices").Where(&condition).First(&c).Error; err != nil {
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
