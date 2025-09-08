package service

import (
	"context"
	"errors"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/auth/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	uRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg"
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

func (s *service) Login(ctx context.Context, u user.User) (*string, error) {
	if u.Nip == "" || u.Password == "" {
		return nil, errors.New("Invalid credentials!")
	}

	user, err := s.repo.GetUser(ctx, u)

	if user == nil || err != nil {
		return nil, errors.New("Invalid credentials!")
	}

	isValid := pkg.CompareHashPassword(u.Password, user.Password)

	if !isValid {
		return nil, errors.New("Invalid credentials!")
	}

	token, err := pkg.CreateJWT(user.ID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *service) Register(ctx context.Context, u user.User) (*string, error) {
	if u.Nip == "" || u.Password == "" || u.Name == "" {
		return nil, errors.New("Invalid request!")
	}

	u.RoleID = 3 // Santri

	hashedPassword, err := pkg.HashPassword(u.Password)

	if err != nil {
		return nil, err
	}

	u.Password = hashedPassword

	user, err := s.userRepo.CreateUser(ctx, u)

	if err != nil {
		return nil, err
	}

	token, err := pkg.CreateJWT(user.ID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
