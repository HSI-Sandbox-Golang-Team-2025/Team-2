package repository

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	"gorm.io/gorm"
)

type userTrackRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserTrackRepository {
	return &userTrackRepository{DB: db}
}

func (r *userTrackRepository) Create(userTrack *user_track.UserTrack) error {
	return r.DB.Create(userTrack).Error
}

func (r *userTrackRepository) GetByID(id uint) (*user_track.UserTrack, error) {
	var ut user_track.UserTrack
	if err := r.DB.First(&ut, id).Error; err != nil {
		return nil, err
	}
	return &ut, nil
}

func (r *userTrackRepository) Update(userTrack *user_track.UserTrack) error {
	return r.DB.Save(userTrack).Error
}

func (r *userTrackRepository) Delete(id uint) error {
	return r.DB.Delete(&user_track.UserTrack{}, id).Error
}

func (r *userTrackRepository) List() ([]user_track.UserTrack, error) {
	var uts []user_track.UserTrack
	if err := r.DB.Find(&uts).Error; err != nil {
		return nil, err
	}
	return uts, nil
}
