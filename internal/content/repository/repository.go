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

func (r *repository) CreateContent(ctx context.Context, c content.Content) (*content.Content, error) {
	if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *repository) GetContentByID(ctx context.Context, id int64) (*content.Content, error) {
	var c content.Content
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *repository) GetAllContent(ctx context.Context) ([]content.Content, error) {
	var contentList []content.Content
	if err := r.db.WithContext(ctx).Find(&contentList).Error; err != nil {
		return nil, err
	}
	return contentList, nil
}

func (r *repository) UpdateContent(ctx context.Context, c content.Content) error {
	return r.db.WithContext(ctx).Save(&c).Error
}

func (r *repository) DeleteContent(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&content.Content{}, id).Error
}
