package service

import (
	"context"
	"slices"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record"
	userPracticeRecordRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	userPracticeRepo       repository.Repository
	userPracticeRecordRepo userPracticeRecordRepository.Repository
}

func NewService(
	userPracticeRepo repository.Repository,
	userPracticeRecordRepo userPracticeRecordRepository.Repository,
) Service {
	return &service{
		userPracticeRepo:       userPracticeRepo,
		userPracticeRecordRepo: userPracticeRecordRepo,
	}
}

func (s *service) StartUserPractice(
	ctx context.Context,
	paramId string,
) (*user_practice.UserPractice, error) {
	userPractice := user_practice.UserPractice{}

	id, _ := strconv.Atoi(paramId)

	condition := repository.GetUserPracticeCondition{
		ID:     uint(id),
		Status: user_practice.Opened,
	}

	err := s.userPracticeRepo.GetUserPractice(
		ctx,
		&userPractice,
		&condition,
	)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "User practice not found!")
	}

	userPractice.Status = user_practice.InProgress

	err = s.userPracticeRepo.UpdateUserPractice(ctx, &userPractice)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userPractice, nil
}

func (s *service) GetUserPractices(
	ctx context.Context,
	queries map[string]string,
) (*[]user_practice.UserPractice, error) {
	userPractices := []user_practice.UserPractice{}

	err := s.userPracticeRepo.GetUserPractices(ctx, &userPractices, queries)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &userPractices, nil
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
) (*user_practice.UserPractice, error) {
	userPractice := user_practice.UserPractice{}

	id, _ := strconv.Atoi(paramId)

	userPracticeRecordIds := []uint{}

	// CHECK: apabila body.UserPracticeRecords kosong, *nil bikin panic!
	for _, bodyRecord := range *body.UserPracticeRecords {
		userPracticeRecordIds = append(userPracticeRecordIds, bodyRecord.ID)
	}

	condition := repository.GetUserPracticeCondition{
		ID:                    uint(id),
		Status:                user_practice.Submitted,
		UserPracticeRecordIds: userPracticeRecordIds,
	}

	err := s.userPracticeRepo.GetUserPractice(ctx, &userPractice, &condition)

	// CHECK: apabila body.UserPracticeRecords kosong, *nil bikin panic!
	if err != nil || len(*userPractice.UserPracticeRecords) != len(*body.UserPracticeRecords) {
		return nil, fiber.NewError(fiber.StatusNotFound, "User practice not found!")
	}

	userPractice.Status = user_practice.Reviewed
	userPractice.Score = body.Score

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

	return &userPractice, nil
}
