package user_content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
)

type UserContent struct {
	basic_model.BasicModel
	UserID    uint            `json:"userId"`
	ContentID uint            `json:"contentId"`
	Content   content.Content `json:"content"`
}
