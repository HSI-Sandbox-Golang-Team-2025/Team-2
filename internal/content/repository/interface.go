package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
)

type Repository interface {
	GetContents(ctx context.Context, c *[]content.Content, condition *GetContentsCondition) error
	GetContent(ctx context.Context, c *content.Content, paramId string) error
	CreateContent(ctx context.Context, c *content.Content) error

	GetContentByID(ctx context.Context, id int64) (*content.Content, error)
	GetAllContent(ctx context.Context) ([]content.Content, error)
	UpdateContent(ctx context.Context, c content.Content) error
	DeleteContent(ctx context.Context, id int64) error
}
