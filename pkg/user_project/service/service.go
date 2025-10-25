package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	userTrackRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	userProjecteRepo repository.Repository
	contentRepo      contentRepository.Repository
	userTrackRepo    userTrackRepository.Repository
}

func NewService(
	userProjecteRepo repository.Repository,
	contentRepo contentRepository.Repository,
	userTrackRepo userTrackRepository.Repository,
) Service {
	return &service{
		userProjecteRepo: userProjecteRepo,
		contentRepo:      contentRepo,
		userTrackRepo:    userTrackRepo,
	}
}

func (s *service) StartUserProject(
	ctx context.Context,
	body user_project.UserProject,
	user user.User,
) (*user_project.UserProject, error) {
	c := content.Content{}

	contentId := body.ContentID

	getContentCondition := contentRepository.GetContentCondition{
		ID:   contentId,
		Type: content.ContentTypeProject,
	}

	err := s.contentRepo.GetContent(ctx, &c, &getContentCondition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userProject := user_project.UserProject{}

	getUserProjectCondition := repository.GetUserProjectCondition{
		ContentID: contentId,
		StatusNot: user_project.Rejected,
	}

	err = s.userProjecteRepo.GetUserProject(
		ctx,
		&userProject,
		&getUserProjectCondition,
	)

	if err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "User project already started!")
	}

	userProject = user_project.UserProject{
		UserID:    user.ID,
		ContentID: body.ContentID,
		Status:    user_project.InProgress,
	}

	err = s.userProjecteRepo.StartUserProject(ctx, &userProject)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProject, nil
}

func (s *service) GetUserProjects(
	ctx context.Context,
	queries map[string]string,
	user user.User,
) (*[]user_project.UserProject, error) {
	userProjects := []user_project.UserProject{}

	contentId, _ := strconv.Atoi(queries["contentId"])

	condition := repository.GetUserProjectsCondition{
		UserID:    user.ID,
		ContentID: uint(contentId),
	}

	err := s.userProjecteRepo.GetUserProjects(
		ctx,
		&userProjects,
		&condition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProjects, nil
}

func (s *service) SubmitUserProject(
	ctx context.Context,
	body user_project.UserProject,
	paramId string,
	user user.User,
) (*user_project.UserProject, error) {
	userProject := user_project.UserProject{}

	id, _ := strconv.Atoi(paramId)

	condition := repository.GetUserProjectCondition{
		ID:     uint(id),
		UserID: user.ID,
		Status: user_project.InProgress,
	}

	err := s.userProjecteRepo.GetUserProject(
		ctx,
		&userProject,
		&condition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
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

	err := s.userProjecteRepo.GetUserProject(
		ctx,
		&userProject,
		&condition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User practice not found!")
	}

	userProject.Status = user_project.Approved
	userProject.Score = body.Score
	userProject.Comment = body.Comment

	err = s.userProjecteRepo.UpdateUserProject(ctx, &userProject)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	calculateUserTrackAverageScoreCondition := userTrackRepository.CalculateUserTrackAverageScore{
		ID: userProject.UserTrackID,
	}

	err = s.userTrackRepo.CalculateUserTrackAverageScore(
		ctx,
		&user_track.UserTrack{},
		&calculateUserTrackAverageScoreCondition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProject, nil
}
