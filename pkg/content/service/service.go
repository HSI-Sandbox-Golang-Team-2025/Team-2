package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"
	userMaterialRepo "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material/repository"
	userPracticeRepo "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice/repository"
	userTrackRepo "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	contentRepo      repository.Repository
	userMaterialRepo userMaterialRepo.Repository
	userPracticeRepo userPracticeRepo.Repository
	userTrackRepo    userTrackRepo.Repository
}

func NewService(
	contentRepo repository.Repository,
	userMaterialRepo userMaterialRepo.Repository,
	userPracticeRepo userPracticeRepo.Repository,
	userTrackRepo userTrackRepo.Repository,
) Service {
	return &service{
		contentRepo:      contentRepo,
		userMaterialRepo: userMaterialRepo,
		userPracticeRepo: userPracticeRepo,
		userTrackRepo:    userTrackRepo,
	}
}

func (s *service) GetContents(
	ctx context.Context,
	queries map[string]string,
	userAuth *user.User,
) (*[]content.Content, error) {
	reqUserId := userAuth.ID
	reqRoleId := userAuth.RoleID

	trackId, _ := strconv.Atoi(queries["trackId"])
	userId, _ := strconv.Atoi(queries["userId"])

	if reqRoleId != uint(role.Mentor) && reqUserId != uint(userId) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	contents := []content.Content{}

	condition := repository.GetContentsCondition{
		TrackID:  uint(trackId),
		Type:     content.ContentType(queries["type"]),
		UserID:   uint(userId),
		HideBody: true,
	}

	if err := s.contentRepo.GetContents(ctx, &contents, &condition); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &contents, nil
}

func (s *service) GetContent(
	ctx context.Context,
	paramId string,
	userAuth *user.User,
) (*content.Content, error) {
	contentRes := content.Content{}

	id, _ := strconv.Atoi(paramId)

	userId := userAuth.ID

	condition := repository.GetContentCondition{
		ID:     uint(id),
		UserID: userId,
	}

	err := s.contentRepo.GetContent(ctx, &contentRes, &condition)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Content not found!")
	}

	isUserTrackValid := false

	validateUserTrackCondition := userTrackRepo.ValidateUserTrackCondition{
		TrackID: contentRes.TrackID,
		UserID:  userId,
	}

	err = s.userTrackRepo.ValidateUserTrack(
		ctx,
		&isUserTrackValid,
		&validateUserTrackCondition,
	)

	if err != nil || !isUserTrackValid {
		return nil, fiber.NewError(fiber.StatusNotFound, "You need to register this track first!")
	}

	// Check are the last contents completed?
	if contentRes.Order > 0 {
		prevCompletedContents := []content.Content{}

		getPrevCompletedContentsCondition := repository.GetContentsCondition{
			TrackID:       contentRes.TrackID,
			OrderBefore:   contentRes.Order,
			UserID:        userId,
			OnlyCompleted: true,
		}

		err := s.contentRepo.GetContents(
			ctx,
			&prevCompletedContents,
			&getPrevCompletedContentsCondition,
		)

		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		if len(prevCompletedContents) != int(contentRes.Order) {
			return nil, fiber.NewError(fiber.StatusForbidden, "Previous content is not completed yet!")
		}
	}

	// Update UserMaterial.Status to "opened"
	if contentRes.Type == content.ContentTypeMaterial {
		userMaterial := user_material.UserMaterial{}

		getUserMaterialCondition := userMaterialRepo.GetUserMaterialCondition{
			UserID:    userId,
			ContentID: contentRes.ID,
			Status:    user_material.Opened,
		}

		err := s.userMaterialRepo.GetUserMaterial(ctx, &userMaterial, &getUserMaterialCondition)

		if err != nil {
			if err.Error() == "Material not found!" {
				userMaterial.UserID = userId
				userMaterial.ContentID = contentRes.ID

				err = s.userMaterialRepo.OpenUserMaterial(ctx, &userMaterial)

				if err != nil {
					return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
				}

				contentRes.IsCompleted = true
				contentRes.UserMaterial = &userMaterial
			} else {
				return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
		}
	}

	contents := []content.Content{}

	getContentsCondition := repository.GetContentsCondition{
		TrackID:  contentRes.TrackID,
		HideBody: true,
	}

	if err := s.contentRepo.GetContents(ctx, &contents, &getContentsCondition); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Update UserTrack.Status to "completed"
	if uint(len(contents)) == contentRes.Order+1 {
		condition := userTrackRepo.CompleteUserTrackCondition{
			TrackID: contentRes.TrackID,
			UserID:  userId,
		}

		err := s.userTrackRepo.CompleteUserTrack(ctx, &condition)

		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

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
