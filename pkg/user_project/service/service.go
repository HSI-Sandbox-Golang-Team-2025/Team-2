package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
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

	userId := user.ID
	contentId := body.ContentID

	getContentCondition := contentRepository.GetContentCondition{
		ID:   contentId,
		Type: content.ContentTypeProject,
	}

	err := s.contentRepo.GetContent(ctx, &c, &getContentCondition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userTrack := user_track.UserTrack{}

	condition := userTrackRepository.GetUserTrackCondition{
		TrackID: c.TrackID,
		UserID:  userId,
	}

	err = s.userTrackRepo.GetUserTrack(ctx, &userTrack, &condition)

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
		UserID:      userId,
		ContentID:   body.ContentID,
		UserTrackID: userTrack.ID,
		Status:      user_project.InProgress,
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
	reqUserId := user.ID
	reqRoleId := user.RoleID

	// if reqRoleId != uint(role.Mentor) && reqUserId != uint(userId) {
	// 	return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	// }

	if reqRoleId != uint(role.Mentor) && queries["userId"] != "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	contentId, _ := strconv.Atoi(queries["contentId"])
	userId, _ := strconv.Atoi(queries["userId"])
	status := queries["status"]

	if reqRoleId != uint(role.Mentor) {
		userId = int(reqUserId)
	}

	userProjects := []user_project.UserProject{}

	condition := repository.GetUserProjectsCondition{
		UserID:    uint(userId),
		ContentID: uint(contentId),
		Status:    user_project.UserProjectStatus(status),
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

	// Update the average score with the new score
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

	// Update UserTrack.Status to "completed"
	completeUserTrackCondition := userTrackRepository.CompleteUserTrackCondition{
		ID: userProject.UserTrackID,
	}

	err = s.userTrackRepo.CompleteUserTrack(ctx, &completeUserTrackCondition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userProject, nil
}
