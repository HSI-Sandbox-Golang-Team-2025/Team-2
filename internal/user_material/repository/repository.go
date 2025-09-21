package repository

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_materials"
	"gorm.io/gorm"
)

type userMaterialRepository struct {
	DB *gorm.DB
}

func NewUserMaterialRepository(db *gorm.DB) UserMaterialRepository {
	return &userMaterialRepository{DB: db}
}

func (r *userMaterialRepository) Create(userMaterial *user_materials.UserMaterial) error {
	return r.DB.Create(userMaterial).Error
}

func (r *userMaterialRepository) GetByID(id uint) (*user_materials.UserMaterial, error) {
	var um user_materials.UserMaterial
	if err := r.DB.First(&um, id).Error; err != nil {
		return nil, err
	}
	return &um, nil
}

func (r *userMaterialRepository) Update(userMaterial *user_materials.UserMaterial) error {
	return r.DB.Save(userMaterial).Error
}

func (r *userMaterialRepository) Delete(id uint) error {
	return r.DB.Delete(&user_materials.UserMaterial{}, id).Error
}

func (r *userMaterialRepository) List() ([]user_materials.UserMaterial, error) {
	var ums []user_materials.UserMaterial
	if err := r.DB.Find(&ums).Error; err != nil {
		return nil, err
	}
	return ums, nil
}
