package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
)

type Service interface {
	GetUserContents(
		ctx context.Context,
		paramUserId string,
		queries map[string]string,
		userAuth *user.User,
	) (*[]content.Content, error)
}
