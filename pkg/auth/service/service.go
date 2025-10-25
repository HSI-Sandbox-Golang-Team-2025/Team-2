package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/lib"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/auth/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	uRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	repo     repository.Repository
	userRepo uRepository.Repository
}

func NewService(authRepo repository.Repository, userRepo uRepository.Repository) Service {
	return &service{
		repo:     authRepo,
		userRepo: userRepo,
	}
}

func (s *service) Login(ctx context.Context, body user.User) (*string, error) {
	if body.Nip == "" || body.Password == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid credentials!")
	}

	u := user.User{}

	condition := uRepository.FindUserCondition{}
	condition.Nip = body.Nip

	err := s.userRepo.GetUser(ctx, &u, &condition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid credentials!")
	}

	isValid := lib.CompareHashPassword(body.Password, u.Password)

	if !isValid {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid credentials!")
	}

	token, err := lib.CreateJWT(u.ID)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &token, nil
}

func (s *service) Register(ctx context.Context, body user.User) (*string, error) {
	if body.Nip == "" || body.Password == "" || body.Name == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request!")
	}

	u := user.User{}

	u.Nip = body.Nip
	u.Name = body.Name
	u.RoleID = uint(role.Santri)

	hashedPassword, err := lib.HashPassword(body.Password)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	u.Password = hashedPassword

	err = s.userRepo.CreateUser(ctx, &u)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	token, err := lib.CreateJWT(u.ID)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &token, nil
}
