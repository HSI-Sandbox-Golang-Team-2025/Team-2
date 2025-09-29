package repository

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
)

type Repository interface {
	StartUserProject(
		ctx context.Context,
		userProject *user_project.UserProject,
	) error
	GetUserProjects(
		ctx context.Context,
		userProjects *[]user_project.UserProject,
		queries map[string]string,
	) error
	GetUserProject(
		ctx context.Context,
		userProject *user_project.UserProject,
		condition *GetUserProjectCondition,
	) error
	UpdateUserProject(
		ctx context.Context,
		userProject *user_project.UserProject,
	) error
}
