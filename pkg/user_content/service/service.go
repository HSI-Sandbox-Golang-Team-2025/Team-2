package service

import (
	"context"
	"strconv"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
	userTrackRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track/repository"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	contentRepo   repository.Repository
	userTrackRepo userTrackRepository.Repository
}

func NewService(
	contentRepo repository.Repository,
	userTrackRepo userTrackRepository.Repository,
) Service {
	return &service{
		contentRepo:   contentRepo,
		userTrackRepo: userTrackRepo,
	}
}

func (s *service) GetUserContents(
	ctx context.Context,
	paramUserId string,
	queries map[string]string,
	userAuth *user.User,
) (*[]content.Content, error) {
	reqUserId := userAuth.ID
	reqRoleId := userAuth.RoleID

	trackId, _ := strconv.Atoi(queries["trackId"])
	userId, _ := strconv.Atoi(paramUserId)

	if trackId == 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Missing query: trackId")
	}

	if userId == 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Missing param: userId")
	}

	if reqRoleId != uint(role.Mentor) && reqUserId != uint(userId) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not authorized!")
	}

	getUserTrackCondition := userTrackRepository.GetUserTrackCondition{
		TrackID: uint(trackId),
		UserID:  uint(userId),
	}

	err := s.userTrackRepo.GetUserTrack(
		ctx,
		&user_track.UserTrack{},
		&getUserTrackCondition,
	)

	if err != nil {
		if err.Error() == "User track not found!" {
			return nil, fiber.NewError(fiber.StatusNotFound, "User is not registered on this track!")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	contents := []content.Content{}

	condition := repository.GetContentsCondition{
		TrackID: uint(trackId),
		Type:    content.ContentType(queries["type"]),
		UserID:  uint(userId),
		// HideBody: true,
	}

	if err := s.contentRepo.GetContents(ctx, &contents, &condition); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &contents, nil
}
