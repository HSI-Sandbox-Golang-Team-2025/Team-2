package service

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_materials"

// Interface untuk service UserMaterials
type UserMaterialService interface {
	Create(userMaterial *user_materials.UserMaterial) error
	GetByID(id uint) (*user_materials.UserMaterial, error)
	Update(userMaterial *user_materials.UserMaterial) error
	Delete(id uint) error
	List() ([]user_materials.UserMaterial, error)
}
