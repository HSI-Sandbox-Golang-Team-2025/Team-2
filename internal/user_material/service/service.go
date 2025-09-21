package service

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_materials"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_materials/repository"
)

type userMaterialService struct {
	repo repository.UserMaterialRepository
}

func NewUserMaterialService(repo repository.UserMaterialRepository) UserMaterialService {
	return &userMaterialService{repo: repo}
}

func (s *userMaterialService) Create(userMaterial *user_materials.UserMaterial) error {
	return s.repo.Create(userMaterial)
}

func (s *userMaterialService) GetByID(id uint) (*user_materials.UserMaterial, error) {
	return s.repo.GetByID(id)
}

func (s *userMaterialService) Update(userMaterial *user_materials.UserMaterial) error {
	return s.repo.Update(userMaterial)
}

func (s *userMaterialService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userMaterialService) List() ([]user_materials.UserMaterial, error) {
	return s.repo.List()
}
