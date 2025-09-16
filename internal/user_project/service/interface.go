package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_project"
)

type Service interface {
	StartUserProject(
		ctx context.Context,
		body user_project.UserProject,
	) (*user_project.UserProject, error)
	GetUserProjects(
		ctx context.Context,
		queries map[string]string,
	) (*[]user_project.UserProject, error)
	SubmitUserProject(
		ctx context.Context,
		body user_project.UserProject,
		paramId string,
	) (*user_project.UserProject, error)
	ReviewUserProject(
		ctx context.Context,
		body user_project.UserProject,
		paramId string,
	) (*user_project.UserProject, error)
}
