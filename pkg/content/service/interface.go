package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Service interface {
	GetContents(
		ctx context.Context,
		queries map[string]string,
		userAuth *user.User,
	) (*[]content.Content, error)
	GetContent(
		ctx context.Context,
		paramId string,
		userAuth *user.User,
	) (*content.Content, error)

	CreateContent(ctx context.Context, c content.Content) (*content.Content, error)
	UpdateContent(ctx context.Context, c content.Content) error
	DeleteContent(ctx context.Context, id int64) error
}
