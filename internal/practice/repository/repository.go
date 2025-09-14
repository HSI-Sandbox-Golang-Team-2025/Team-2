package repository

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreatePractice(ctx context.Context, p *practice.Practice) error {
	if err := r.db.WithContext(ctx).Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetPractices(ctx context.Context, p *[]practice.Practice, queries map[string]string) error {
	trackId, _ := strconv.Atoi(queries["trackId"])

	condition := practice.Practice{
		TrackID: uint(trackId),
	}

	if err := r.db.WithContext(ctx).Preload("Questions").Preload("Questions.AnswerChoices").Where(&condition).Find(&p).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPractice(ctx context.Context, p *practice.Practice, paramId string) error {
	id, _ := strconv.Atoi(paramId)

	condition := practice.Practice{}
	condition.ID = uint(id)

	if err := r.db.WithContext(ctx).Preload("Questions").Preload("Questions.AnswerChoices").Where(&condition).First(&p).Error; err != nil {
		return err
	}

	return nil
}
