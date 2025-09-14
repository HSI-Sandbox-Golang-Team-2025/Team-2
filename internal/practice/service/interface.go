package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
)

type Service interface {
	CreatePractice(ctx context.Context, c content.Content) (*content.Content, error)
}
