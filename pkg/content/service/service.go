package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	userPracticeRepo "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	contentRepo      repository.Repository
	userPracticeRepo userPracticeRepo.Repository
}

func NewService(contentRepo repository.Repository, userPracticeRepo userPracticeRepo.Repository) Service {
	return &service{
		contentRepo:      contentRepo,
		userPracticeRepo: userPracticeRepo,
	}
}

func (s *service) GetContents(ctx context.Context, queries map[string]string, user user.User) (*[]content.Content, error) {
	contents := []content.Content{}

	trackId, _ := strconv.Atoi(queries["trackId"])

	userId := user.ID

	condition := repository.GetContentsCondition{
		TrackID:  uint(trackId),
		Type:     content.ContentType(queries["type"]),
		UserID:   userId,
		HideBody: true,
	}

	if err := s.contentRepo.GetContents(ctx, &contents, &condition); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &contents, nil
}

func (s *service) GetContent(ctx context.Context, paramId string, user user.User) (*content.Content, error) {
	contentRes := content.Content{}

	id, _ := strconv.Atoi(paramId)

	userId := user.ID

	condition := repository.GetContentCondition{
		ID:     uint(id),
		UserID: userId,
	}

	if err := s.contentRepo.GetContent(ctx, &contentRes, &condition); err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Content not found!")
	}

	if contentRes.Order > 1 {
		prevCompletedContents := []content.Content{}

		getPrevCompletedContentsCondition := repository.GetContentsCondition{
			TrackID:       contentRes.TrackID,
			OrderBefore:   contentRes.Order,
			UserID:        userId,
			OnlyCompleted: true,
		}

		if err := s.contentRepo.GetContents(ctx, &prevCompletedContents, &getPrevCompletedContentsCondition); err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		// Is last contents completed?
		if len(prevCompletedContents) != int(contentRes.Order-1) {
			return nil, fiber.NewError(fiber.StatusForbidden, "Previous content is not completed yet!")
		}
	}

	// TODO: Create record to the user_contents

	// if contentRes.Type == content.ContentTypePractice {
	// 	userPractice := user_practice.UserPractice{}

	// 	getUserPracticeCondition := userPracticeRepo.GetUserPracticeCondition{
	// 		UserID:    userId,
	// 		ContentID: contentRes.ID,
	// 		// Status:    user_practice.Opened,
	// 	}

	// 	err := s.userPracticeRepo.GetUserPractice(ctx, &userPractice, &getUserPracticeCondition)

	// 	if err != nil && err.Error() == "record not found" {
	// 		userPractice.UserID = userId
	// 		userPractice.ContentID = contentRes.ID
	// 		// userPractice.Status = user_practice.Opened

	// 		// CHECK: Kalo dia ujian berkali-kali gimana? apakah masih ada status opened?
	// 		// s.userPracticeRepo.OpenUserPractice(ctx, &userPractice)
	// 	}
	// }

	return &contentRes, nil
}

func (s *service) CreateContent(ctx context.Context, c content.Content) (*content.Content, error) {
	c.Type = content.ContentTypeMaterial

	if err := s.contentRepo.CreateContent(ctx, &c); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &c, nil
}

func (s *service) UpdateContent(ctx context.Context, c content.Content) error {
	return s.contentRepo.UpdateContent(ctx, c)
}

func (s *service) DeleteContent(ctx context.Context, id int64) error {
	return s.contentRepo.DeleteContent(ctx, id)
}
