package service

import (
	"context"
	"slices"
	"strconv"

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
}

func NewService(
	userPracticeRepo repository.Repository,
	userPracticeRecordRepo userPracticeRecordRepository.Repository,
	userTrackRepo userTrackRepository.Repository,
) Service {
	return &service{
		userPracticeRepo:       userPracticeRepo,
		userPracticeRecordRepo: userPracticeRecordRepo,
		userTrackRepo:          userTrackRepo,
	}
}

func (s *service) GetUserPractices(
	ctx context.Context,
	queries map[string]string,
	user *user.User,
) (*[]user_practice.UserPractice, error) {
	reqUserId := user.ID
	reqRoleId := user.RoleID

	contentId, _ := strconv.Atoi(queries["contentId"])
	userId, _ := strconv.Atoi(queries["userId"])

	if reqRoleId != uint(role.Mentor) && reqUserId != uint(userId) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	userPractices := []user_practice.UserPractice{}

	condition := repository.GetUserPracticesCondition{
		UserID:    uint(userId),
		ContentID: uint(contentId),
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
	userPractice := user_practice.UserPractice{}

	userPractice.UserID = user.ID
	userPractice.ContentID = body.ContentID
	userPractice.Status = user_practice.InProgress

	err := s.userPracticeRepo.StartUserPractice(ctx, &userPractice)

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
