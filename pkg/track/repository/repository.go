package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateTrack(
	ctx context.Context,
	t *track.Track,
) error {
	findTrackCondition := FindTrackCondition{Track: *t}

	if err := r.GetTrack(ctx, t, &findTrackCondition); err == nil {
		return errors.New("Track is already registered")
	}

	if err := r.db.WithContext(ctx).Create(&t).Error; err != nil {
		return err
	}

	return nil
}

type FindTrackCondition struct {
	track.Track
}

func (r *repository) GetTrack(
	ctx context.Context,
	t *track.Track,
	condition *FindTrackCondition,
) error {
	db := r.db.WithContext(ctx)

	if condition.ID != 0 {
		db = db.Where("id = ?", condition.ID)
	}

	if condition.Name != "" {
		db = db.Where("name = ?", condition.Name)
	}

	err := db.First(&t).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetTrackByID(ctx context.Context, id int64) (*track.Track, error) {
	var t track.Track
	if err := r.db.WithContext(ctx).First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) GetAllTrack(ctx context.Context) ([]track.Track, error) {
	var trackList []track.Track
	if err := r.db.WithContext(ctx).Find(&trackList).Error; err != nil {
		return nil, err
	}
	return trackList, nil
}

func (r *repository) UpdateTrack(ctx context.Context, t track.Track) error {
	return r.db.WithContext(ctx).Save(&t).Error
}

func (r *repository) DeleteTrack(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&track.Track{}, id).Error
}
