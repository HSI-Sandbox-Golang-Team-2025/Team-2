package service

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_material"

// Interface untuk service UserMaterials
type UserMaterialService interface {
	Create(userMaterial *user_material.UserMaterial) error
	GetByID(id uint) (*user_material.UserMaterial, error)
	Update(userMaterial *user_material.UserMaterial) error
	Delete(id uint) error
	List() ([]user_material.UserMaterial, error)
}
