package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	userRepo repository.Repository
}

func NewService(
	userRepo repository.Repository,
) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetUsers(
	ctx context.Context,
	queries map[string]string,
	userAuth *user.User,
) (*[]user.User, error) {
	users := []user.User{}

	getUsersCondtion := repository.FindUsersCondition{}

	err := s.userRepo.GetUsers(ctx, &users, &getUsersCondtion)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &users, nil
}
