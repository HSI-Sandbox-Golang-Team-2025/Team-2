package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, u *user.User) error {
	findUserCondition := FindUserCondition{
		Nip: u.Nip,
	}

	if err := r.GetUser(ctx, u, &findUserCondition); err == nil {
		return errors.New("NIP is already registered")
	}

	if err := r.db.WithContext(ctx).Create(&u).Error; err != nil {
		return err
	}

	return nil
}

type FindUsersCondition struct {
	Nip string
}

func (r *repository) GetUsers(
	ctx context.Context,
	users *[]user.User,
	condition *FindUsersCondition,
) error {
	db := r.db.WithContext(ctx).
		Preload("UserTracks").
		Preload("UserTracks.Track").
		// Preload("UserTracks.Track.Contents", func(db *gorm.DB) *gorm.DB {
		// 	return db.Order("contents.order ASC")
		// }).
		// Preload("UserTracks.Track.Contents.UserMaterial").
		// Preload("UserTracks.Track.Contents.UserPractices").
		// Preload("UserTracks.Track.Contents.UserProjects").
		Order("id")

	if condition.Nip != "" {
		db = db.Where("nip = ?", condition.Nip)
	}

	return db.Find(&users).Error
}

type FindUserCondition struct {
	Nip string
}

func (r *repository) GetUser(
	ctx context.Context,
	usr *user.User,
	condition *FindUserCondition,
) error {
	users := []user.User{}

	getUsersCondition := FindUsersCondition{
		Nip: condition.Nip,
	}

	err := r.GetUsers(ctx, &users, &getUsersCondition)

	if err != nil {
		return err
	}

	*usr = users[0]

	return nil
}
