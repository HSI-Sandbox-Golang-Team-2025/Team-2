package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateMaterial(ctx context.Context, m material.Material) (*material.Material, error) {
	return s.repository.CreateMaterial(ctx, m)
}

func (s *service) GetMaterialByID(ctx context.Context, id int64) (*material.Material, error) {
	return s.repository.GetMaterialByID(ctx, id)
}

func (s *service) GetAllMaterial(ctx context.Context) ([]material.Material, error) {
	return s.repository.GetAllMaterial(ctx)
}

func (s *service) UpdateMaterial(ctx context.Context, m material.Material) error {
	return s.repository.UpdateMaterial(ctx, m)
}

func (s *service) DeleteMaterial(ctx context.Context, id int64) error {
	return s.repository.DeleteMaterial(ctx, id)
}
