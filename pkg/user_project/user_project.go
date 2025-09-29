package user_project

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project_media"
)

type UserProject struct {
	basic_model.BasicModel
	UserID    uint                                   `json:"userId"`
	User      *user.User                             `json:"user"`
	ContentID uint                                   `json:"contentId"`
	Content   *content.Content                       `json:"content"`
	Status    UserProjectStatus                      `json:"status"`
	Url       *string                                `json:"url"`
	Score     float32                                `json:"score"`
	Comment   *string                                `json:"comment"`
	Medias    *[]user_project_media.UserProjectMedia `json:"medias"`
}

type UserProjectStatus string

const (
	InProgress UserProjectStatus = "in progress"
	Submitted  UserProjectStatus = "submitted"
	Rejected   UserProjectStatus = "rejected"
	Approved   UserProjectStatus = "approved"
)
