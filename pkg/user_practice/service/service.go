package service

import (
	"context"
	"slices"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	contentRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record"
	userPracticeRecordRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	userTrackRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	userPracticeRepo       repository.Repository
	userPracticeRecordRepo userPracticeRecordRepository.Repository
	userTrackRepo          userTrackRepository.Repository
	contentRepo            contentRepository.Repository
}

func NewService(
	userPracticeRepo repository.Repository,
	userPracticeRecordRepo userPracticeRecordRepository.Repository,
	userTrackRepo userTrackRepository.Repository,
	contentRepo contentRepository.Repository,
) Service {
	return &service{
		userPracticeRepo:       userPracticeRepo,
		userPracticeRecordRepo: userPracticeRecordRepo,
		userTrackRepo:          userTrackRepo,
		contentRepo:            contentRepo,
	}
}

func (s *service) GetUserPractices(
	ctx context.Context,
	queries map[string]string,
	user *user.User,
) (*[]user_practice.UserPractice, error) {
	reqUserId := user.ID
	reqRoleId := user.RoleID

	if reqRoleId != uint(role.Mentor) && queries["userId"] != "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	contentId, _ := strconv.Atoi(queries["contentId"])
	userId, _ := strconv.Atoi(queries["userId"])
	status := queries["status"]

	if reqRoleId != uint(role.Mentor) {
		userId = int(reqUserId)
	}

	userPractices := []user_practice.UserPractice{}

	condition := repository.GetUserPracticesCondition{
		UserID:    uint(userId),
		ContentID: uint(contentId),
		Status:    user_practice.UserPracticeStatus(status),
	}

	err := s.userPracticeRepo.GetUserPractices(ctx, &userPractices, &condition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userPractices, nil
}

func (s *service) StartUserPractice(
	ctx context.Context,
	body user_practice.UserPractice,
	user user.User,
) (*user_practice.UserPractice, error) {
	userId := user.ID

	contentId := body.ContentID

	getContentCondition := contentRepository.GetContentCondition{
		ID:   contentId,
		Type: content.ContentTypePractice,
	}

	content := content.Content{}

	err := s.contentRepo.GetContent(ctx, &content, &getContentCondition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userTrack := user_track.UserTrack{}

	condition := userTrackRepository.GetUserTrackCondition{
		TrackID: content.TrackID,
		UserID:  userId,
	}

	err = s.userTrackRepo.GetUserTrack(ctx, &userTrack, &condition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userPractice := user_practice.UserPractice{}

	userPractice.UserID = userId
	userPractice.ContentID = body.ContentID
	userPractice.UserTrackID = userTrack.ID
	userPractice.Status = user_practice.InProgress

	err = s.userPracticeRepo.StartUserPractice(ctx, &userPractice)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userPractice, nil
}

func (s *service) SubmitUserPractice(
	ctx context.Context,
	up user_practice.UserPractice,
	paramId string,
) (*user_practice.UserPractice, error) {
	userPractice := user_practice.UserPractice{}

	id, _ := strconv.Atoi(paramId)

	condition := repository.GetUserPracticeCondition{
		ID:     uint(id),
		Status: user_practice.InProgress,
	}

	err := s.userPracticeRepo.GetUserPractice(
		ctx,
		&userPractice,
		&condition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "User practice not found!")
	}

	userPractice.Status = user_practice.Submitted
	userPractice.UserPracticeRecords = up.UserPracticeRecords

	err = s.userPracticeRepo.UpdateUserPractice(ctx, &userPractice)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userPractice, nil
}

func (s *service) ReviewUserPractice(
	ctx context.Context,
	body user_practice.UserPractice,
	paramId string,
	user user.User,
) (*user_practice.UserPractice, error) {
	userPractice := user_practice.UserPractice{}

	id, _ := strconv.Atoi(paramId)

	userPracticeRecordIds := []uint{}

	// TODO: apabila body.UserPracticeRecords kosong, *nil bikin panic!
	for _, bodyRecord := range *body.UserPracticeRecords {
		userPracticeRecordIds = append(userPracticeRecordIds, bodyRecord.ID)
	}

	condition := repository.GetUserPracticeCondition{
		ID:                    uint(id),
		Status:                user_practice.Submitted,
		UserPracticeRecordIds: userPracticeRecordIds,
	}

	err := s.userPracticeRepo.GetUserPractice(ctx, &userPractice, &condition)

	// TODO: apabila body.UserPracticeRecords kosong, *nil bikin panic!
	if err != nil || len(*userPractice.UserPracticeRecords) != len(*body.UserPracticeRecords) {
		return nil, fiber.NewError(fiber.StatusNotFound, "User practice not found!")
	}

	userPractice.Status = user_practice.Reviewed
	userPractice.Score = body.Score
	userPractice.Comment = body.Comment

	err = s.userPracticeRepo.UpdateUserPractice(ctx, &userPractice)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	slices.SortFunc(
		*body.UserPracticeRecords,
		func(i, j user_practice_record.UserPracticeRecord) int {
			return int(i.ID - j.ID) // Ascending sort
		},
	)

	for i, bodyRecord := range *body.UserPracticeRecords {
		userPracticeRecord := *userPractice.UserPracticeRecords
		userPracticeRecord[i].IsCorrect = bodyRecord.IsCorrect
	}

	err = s.userPracticeRecordRepo.UpdateUserPracticeRecords(
		ctx,
		userPractice.UserPracticeRecords,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	calculateUserTrackAverageScoreCondition := userTrackRepository.CalculateUserTrackAverageScore{
		ID: userPractice.UserTrackID,
	}

	err = s.userTrackRepo.CalculateUserTrackAverageScore(
		ctx,
		&user_track.UserTrack{},
		&calculateUserTrackAverageScoreCondition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userPractice, nil
}
