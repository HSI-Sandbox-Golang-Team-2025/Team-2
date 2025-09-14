package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
)

type Repository interface {
	CreatePractice(ctx context.Context, p *practice.Practice) error
	GetPractices(ctx context.Context, p *[]practice.Practice, queries map[string]string) error
	GetPractice(ctx context.Context, p *practice.Practice, paramId string) error
}
