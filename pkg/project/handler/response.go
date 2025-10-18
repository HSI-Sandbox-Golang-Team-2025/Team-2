package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
)

type CreateProjectSuccessResponse struct {
	Message string `json:"message" example:"Create project content success!"`
	Data    Content
}

type CreateProjectErrorResponse struct {
	Message string `json:"message" example:"[Error message]"`
}

type Content struct {
	content.Content
	UserPractices *[]user_project.UserProject `json:"userProjects" swaggertype:"string" example:"null" extensions:"x-order=11"`
}

type ContentType string

const (
	ContentTypeMaterial ContentType = "material"
	ContentTypePractice ContentType = "practice"
	ContentTypeProject  ContentType = "project"
)
