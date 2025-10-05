package handler

import "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/content"

type GetContentsSuccessResponse struct {
	Message string `json:"message" example:"Get content success!"`
	Data    []content.Content
}

type GetContentSuccessResponse struct {
	Message string `json:"message" example:"Get content success!"`
	Data    content.Content
}
