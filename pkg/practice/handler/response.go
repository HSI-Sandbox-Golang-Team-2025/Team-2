package handler

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
)

type CreatePracticeSuccessResponse struct {
	Message string `json:"message" example:"Create practice content success!"`
	Data    Content
}

type CreatePracticeErrorResponse struct {
	Message string `json:"message" example:"[Error message]"`
}

type Content struct {
	content.Content
	UserPractices *[]user_practice.UserPractice `json:"userPractices" swaggertype:"string" example:"null" extensions:"x-order=11"`
}

type ContentType string

const (
	ContentTypeMaterial ContentType = "material"
	ContentTypePractice ContentType = "practice"
	ContentTypeProject  ContentType = "project"
)
