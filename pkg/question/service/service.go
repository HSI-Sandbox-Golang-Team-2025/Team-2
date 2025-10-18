package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
	questionRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	questionRepo questionRepository.Repository
}

func NewService(
	questionRepo questionRepository.Repository,
) Service {
	return &service{
		questionRepo: questionRepo,
	}
}

func (s *service) GetQuestions(
	ctx context.Context,
	queries map[string]string,
	user user.User,
) (*[]question.Question, error) {
	question := []question.Question{}

	// TODO: Cek apakah user elgible untuk lihat question

	contentId, _ := strconv.Atoi(queries["contentId"])

	condition := questionRepository.GetQuestionCondition{}
	condition.ContentID = uint(contentId)

	err := s.questionRepo.GetQuestions(ctx, &question, &condition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &question, nil
}
