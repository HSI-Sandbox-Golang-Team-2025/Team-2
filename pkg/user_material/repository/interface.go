package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"
)

// Interface untuk repository UserMaterials
type Repository interface {
	OpenUserMaterial(
		ctx context.Context,
		userMaterial *user_material.UserMaterial,
	) error
	GetUserMaterials(
		ctx context.Context,
		userMaterials *[]user_material.UserMaterial,
		condition *GetUserMaterialsCondition,
	) error
	GetUserMaterial(
		ctx context.Context,
		userMaterial *user_material.UserMaterial,
		condition *GetUserMaterialCondition,
	) error

	Create(userMaterial *user_material.UserMaterial) error
	GetByID(id uint) (*user_material.UserMaterial, error)
	Update(userMaterial *user_material.UserMaterial) error
	Delete(id uint) error
	List() ([]user_material.UserMaterial, error)
}
