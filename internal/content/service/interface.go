package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
)

type Service interface {
	GetContents(ctx context.Context, queries map[string]string) (*[]content.Content, error)
	GetContent(ctx context.Context, paramId string) (*content.Content, error)
}
