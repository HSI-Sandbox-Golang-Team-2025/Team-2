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

func (r *repository) GetContents(ctx context.Context, c *[]content.Content, queries map[string]string) error {
	trackId, _ := strconv.Atoi(queries["trackId"])

	condition := content.Content{}
	condition.TrackID = uint(trackId)

	if err := r.db.WithContext(ctx).Preload("Questions").Preload("Questions.AnswerChoices").Where(&condition).Find(&c).Error; err != nil {
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

func (r *repository) CreateContent(ctx context.Context, c *content.Content) error {
	if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
		return err
	}
	return nil
}
