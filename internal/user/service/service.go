package service

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}
