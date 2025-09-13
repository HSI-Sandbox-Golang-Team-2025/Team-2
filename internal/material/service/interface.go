package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/material"
)

type Service interface {
	CreateMaterial(ctx context.Context, m material.Material) (*material.Material, error)
	GetMaterialByID(ctx context.Context, id int64) (*material.Material, error)
	GetAllMaterial(ctx context.Context) ([]material.Material, error)
	UpdateMaterial(ctx context.Context, m material.Material) error
	DeleteMaterial(ctx context.Context, id int64) error
}
