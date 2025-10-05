package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/material"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateMaterial(ctx context.Context, m material.Material) (*material.Material, error) {
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) GetMaterialByID(ctx context.Context, id int64) (*material.Material, error) {
	var m material.Material
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) GetAllMaterial(ctx context.Context) ([]material.Material, error) {
	var materialList []material.Material
	if err := r.db.WithContext(ctx).Find(&materialList).Error; err != nil {
		return nil, err
	}
	return materialList, nil
}

func (r *repository) UpdateMaterial(ctx context.Context, m material.Material) error {
	return r.db.WithContext(ctx).Save(&m).Error
}

func (r *repository) DeleteMaterial(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&material.Material{}, id).Error
}
