package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(
	ctx context.Context,
	userTrack *user_track.UserTrack,
) error {
	isRegistered := false

	validateUserTrackCondition := &ValidateUserTrackCondition{
		UserID:  userTrack.UserID,
		TrackID: userTrack.TrackID,
	}

	err := r.ValidateUserTrack(
		ctx,
		&isRegistered,
		validateUserTrackCondition,
	)

	if err != nil {
		return err
	}

	if isRegistered {
		return errors.New("You can't register for the same track twice!")
	}

	return r.db.WithContext(ctx).
		Create(userTrack).
		Error
}

func (r *repository) GetByID(id uint) (*user_track.UserTrack, error) {
	var ut user_track.UserTrack
	if err := r.db.First(&ut, id).Error; err != nil {
		return nil, err
	}
	return &ut, nil
}

func (r *repository) Update(userTrack *user_track.UserTrack) error {
	return r.db.Save(userTrack).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&user_track.UserTrack{}, id).Error
}

func (r *repository) List() ([]user_track.UserTrack, error) {
	var uts []user_track.UserTrack
	if err := r.db.Find(&uts).Error; err != nil {
		return nil, err
	}
	return uts, nil
}

type GetUserTracksCondition struct {
	TrackID uint
	UserID  uint
	Limit   uint
}

func (r *repository) GetUserTracks(
	ctx context.Context,
	userTrack *[]user_track.UserTrack,
	condition *GetUserTracksCondition,
) error {
	db := r.db.WithContext(ctx)

	if condition.TrackID != 0 {
		db = db.Where("track_id = ?", condition.TrackID)
	}

	if condition.UserID != 0 {
		db = db.Where("user_id = ?", condition.UserID)
	}

	if condition.Limit != 0 {
		db = db.Limit(int(condition.Limit))
	}

	return db.
		Find(&userTrack).
		Error
}

type ValidateUserTrackCondition struct {
	TrackID uint
	UserID  uint
}

func (r *repository) ValidateUserTrack(
	ctx context.Context,
	isValid *bool,
	condition *ValidateUserTrackCondition,
) error {
	userTracks := []user_track.UserTrack{}

	getUserTracksCondition := GetUserTracksCondition{
		TrackID: condition.TrackID,
		UserID:  condition.UserID,
		Limit:   1,
	}

	err := r.GetUserTracks(
		ctx,
		&userTracks,
		&getUserTracksCondition,
	)

	if err != nil {
		return err
	}

	*isValid = len(userTracks) > 0

	return nil
}
