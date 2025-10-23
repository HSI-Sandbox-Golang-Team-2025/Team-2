package repository

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) OpenUserMaterial(
	ctx context.Context,
	userMaterial *user_material.UserMaterial,
) error {
	userMaterial.Status = user_material.Opened

	if err := r.db.WithContext(ctx).Create(&userMaterial).Error; err != nil {
		return err
	}

	return nil
}

type GetUserMaterialsCondition struct {
	ID        uint
	UserID    uint
	ContentID uint
	Status    user_material.UserMaterialStatus
	Limit     uint
}

func (r *repository) GetUserMaterials(
	ctx context.Context,
	userMaterials *[]user_material.UserMaterial,
	condition *GetUserMaterialsCondition,
) error {
	db := r.db.
		WithContext(ctx)

	if condition.ID != 0 {
		db = db.Where("id = ?", condition.ID)
	}

	if condition.UserID != 0 {
		db = db.Where("user_id = ?", condition.UserID)
	}

	if condition.ContentID != 0 {
		db = db.Where("content_id = ?", condition.ContentID)
	}

	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if err := db.Find(&userMaterials).Error; err != nil {
		return err
	}

	return nil
}

type GetUserMaterialCondition struct {
	ID        uint
	UserID    uint
	ContentID uint
	Status    user_material.UserMaterialStatus
}

func (r *repository) GetUserMaterial(
	ctx context.Context,
	userMaterial *user_material.UserMaterial,
	condition *GetUserMaterialCondition,
) error {
	userMaterials := []user_material.UserMaterial{}

	getUserMaterialsCondition := GetUserMaterialsCondition{
		ID:        condition.ID,
		UserID:    condition.UserID,
		ContentID: condition.ContentID,
		Status:    condition.Status,
		Limit:     1,
	}

	err := r.GetUserMaterials(ctx, &userMaterials, &getUserMaterialsCondition)

	if err != nil || len(userMaterials) < 1 {
		return errors.New("Material not found!")
	}

	*userMaterial = userMaterials[0]

	return nil
}

func (r *repository) Create(userMaterial *user_material.UserMaterial) error {
	return r.db.Create(userMaterial).Error
}

func (r *repository) GetByID(id uint) (*user_material.UserMaterial, error) {
	var um user_material.UserMaterial
	if err := r.db.First(&um, id).Error; err != nil {
		return nil, err
	}
	return &um, nil
}

func (r *repository) Update(userMaterial *user_material.UserMaterial) error {
	return r.db.Save(userMaterial).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&user_material.UserMaterial{}, id).Error
}

func (r *repository) List() ([]user_material.UserMaterial, error) {
	var ums []user_material.UserMaterial
	if err := r.db.Find(&ums).Error; err != nil {
		return nil, err
	}
	return ums, nil
}
