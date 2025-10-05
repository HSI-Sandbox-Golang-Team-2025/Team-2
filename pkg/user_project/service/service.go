package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	userProjecteRepo repository.Repository
}

func NewService(
	userProjecteRepo repository.Repository,
) Service {
	return &service{
		userProjecteRepo: userProjecteRepo,
	}
}

func (s *service) StartUserProject(
	ctx context.Context,
	body user_project.UserProject,
) (*user_project.UserProject, error) {
	userProject := user_project.UserProject{
		UserID:    body.UserID,
		ContentID: body.ContentID,
		Status:    user_project.InProgress,
	}

	err := s.userProjecteRepo.StartUserProject(ctx, &userProject)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProject, nil
}

func (s *service) GetUserProjects(
	ctx context.Context,
	queries map[string]string,
) (*[]user_project.UserProject, error) {
	userProjects := []user_project.UserProject{}

	err := s.userProjecteRepo.GetUserProjects(ctx, &userProjects, queries)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProjects, nil
}

func (s *service) SubmitUserProject(
	ctx context.Context,
	body user_project.UserProject,
	paramId string,
) (*user_project.UserProject, error) {
	userProject := user_project.UserProject{}

	id, _ := strconv.Atoi(paramId)

	condition := repository.GetUserProjectCondition{
		ID:     uint(id),
		Status: user_project.InProgress,
	}

	err := s.userProjecteRepo.GetUserProject(
		ctx,
		&userProject,
		&condition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "User practice not found!")
	}

	userProject.Status = user_project.Submitted
	userProject.Medias = body.Medias

	err = s.userProjecteRepo.UpdateUserProject(ctx, &userProject)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProject, nil
}

func (s *service) ReviewUserProject(
	ctx context.Context,
	body user_project.UserProject,
	paramId string,
) (*user_project.UserProject, error) {
	userProject := user_project.UserProject{}

	id, _ := strconv.Atoi(paramId)

	condition := repository.GetUserProjectCondition{
		ID:     uint(id),
		Status: user_project.Submitted,
	}

	err := s.userProjecteRepo.GetUserProject(ctx, &userProject, &condition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User practice not found!")
	}

	userProject.Status = user_project.Approved
	userProject.Score = body.Score

	err = s.userProjecteRepo.UpdateUserProject(ctx, &userProject)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProject, nil
}
