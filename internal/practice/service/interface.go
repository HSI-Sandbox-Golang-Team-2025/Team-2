package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/practice"
)

type Service interface {
	CreatePractice(ctx context.Context, p practice.Practice) (*practice.Practice, error)
	GetPractices(ctx context.Context, queries map[string]string) (*[]practice.Practice, error)
	GetPractice(ctx context.Context, paramId string) (*practice.Practice, error)
}
