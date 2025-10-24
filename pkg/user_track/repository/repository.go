package repository

import (
	"context"
	"errors"
	"fmt"

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

	userTrack.Status = user_track.InProgress

	return r.db.WithContext(ctx).
		Create(userTrack).
		Error
}

type GetUserTracksCondition struct {
	TrackID uint
	UserID  uint
	Status  user_track.UserTrackStatus
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

	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.Limit != 0 {
		db = db.Limit(int(condition.Limit))
	}

	return db.
		Find(&userTrack).
		Error
}

type GetUserTrackCondition struct {
	TrackID uint
	UserID  uint
	Status  user_track.UserTrackStatus
}

func (r *repository) GetUserTrack(
	ctx context.Context,
	userTrack *user_track.UserTrack,
	condition *GetUserTrackCondition,
) error {
	userTracks := []user_track.UserTrack{}

	getUserTracksCondition := GetUserTracksCondition{
		TrackID: condition.TrackID,
		UserID:  condition.UserID,
		Status:  condition.Status,
		Limit:   1,
	}

	err := r.GetUserTracks(ctx, &userTracks, &getUserTracksCondition)

	if err != nil || len(userTracks) < 1 {
		return errors.New("User track not found!")
	}

	*userTrack = userTracks[0]

	return nil
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

type CompleteUserTrackCondition struct {
	TrackID uint
	UserID  uint
}

func (r *repository) CompleteUserTrack(
	ctx context.Context,
	condition *CompleteUserTrackCondition,
) error {
	userTrack := user_track.UserTrack{}

	getUserTrackCondition := GetUserTrackCondition{
		UserID:  condition.UserID,
		TrackID: condition.TrackID,
		Status:  user_track.InProgress,
	}

	err := r.GetUserTrack(
		ctx,
		&userTrack,
		&getUserTrackCondition,
	)

	if err != nil {
		if err.Error() == "User track not found!" {
			return nil
		}
		return err
	}

	userTrack.Status = user_track.Completed

	return r.db.WithContext(ctx).
		Updates(&userTrack).
		Error
}

type CalculateUserTrackAverageScore struct {
	ID      uint
	TrackID uint
	UserID  uint
}

func (r *repository) CalculateUserTrackAverageScore(
	ctx context.Context,
	userTrack *user_track.UserTrack,
	condition *CalculateUserTrackAverageScore,
) error {
	err := r.db.WithContext(ctx).
		Select(`user_tracks.*, COALESCE(AVG(contents.score), 0) AS "average_score"`).
		Joins(
			"LEFT JOIN (?) contents on contents.track_id = user_tracks.id",
			r.db.WithContext(ctx).
				Table("contents").
				Select(`contents.track_id, COALESCE(MAX(user_practices.score), MAX(user_projects.score)) AS "score"`).
				Joins("LEFT JOIN user_practices ON user_practices.content_id = contents.id").
				Joins("LEFT JOIN user_projects ON user_projects.content_id = contents.id").
				Where("contents.type IN ('practice', 'project')").
				Where("COALESCE(user_practices.score, user_projects.score) IS NOT NULL").
				Group("contents.id"),
		).
		Where("user_tracks.id = ?", condition.ID).
		Group("user_tracks.id").
		First(&userTrack).
		Error

	if err != nil {
		return err
	}

	fmt.Println(userTrack.AverageScore)

	return r.UpdateUserTrack(ctx, userTrack)
}

func (r *repository) UpdateUserTrack(
	ctx context.Context,
	userTrack *user_track.UserTrack,
) error {
	if err := r.db.WithContext(ctx).Updates(&userTrack).Error; err != nil {
		return err
	}
	return nil
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
