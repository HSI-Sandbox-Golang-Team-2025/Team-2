package repository

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"

// Interface untuk repository UserMaterials
type UserMaterialRepository interface {
	Create(userMaterial *user_material.UserMaterial) error
	GetByID(id uint) (*user_material.UserMaterial, error)
	Update(userMaterial *user_material.UserMaterial) error
	Delete(id uint) error
	List() ([]user_material.UserMaterial, error)
}
