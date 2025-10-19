package service

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
)

type Service interface {
	StartUserProject(
		ctx context.Context,
		body user_project.UserProject,
		user user.User,
	) (*user_project.UserProject, error)
	GetUserProjects(
		ctx context.Context,
		queries map[string]string,
		user user.User,
	) (*[]user_project.UserProject, error)
	SubmitUserProject(
		ctx context.Context,
		body user_project.UserProject,
		paramId string,
		user user.User,
	) (*user_project.UserProject, error)
	ReviewUserProject(
		ctx context.Context,
		body user_project.UserProject,
		paramId string,
	) (*user_project.UserProject, error)
}
