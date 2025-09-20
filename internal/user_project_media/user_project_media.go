package user_project_media

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/basic_model"
)

type UserProjectMedia struct {
	basic_model.BasicModel
	UserProjectID uint   `json:"userProjectId"`
	Name          string `json:"name"`
	Url           string `json:"url"`
}
