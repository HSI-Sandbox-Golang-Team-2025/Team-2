package user_content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
)

type UserContent struct {
	basic_model.BasicModel
	UserID    uint              `json:"userId"`
	User      user.User         `json:"user"`
	ContentID uint              `json:"contentId"`
	Content   content.Content   `json:"content"`
	Status    UserContentStatus `json:"status"`
}

type UserContentStatus string

const (
	Opened    UserContentStatus = "opened"
	Completed UserContentStatus = "completed"
)
