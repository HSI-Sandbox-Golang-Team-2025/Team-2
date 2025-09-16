package user_practice

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/content"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/user_practice_record"
)

type UserPractice struct {
	basic_model.BasicModel
	UserID              uint                                       `json:"userId"`
	User                *user.User                                 `json:"user"`
	ContentID           uint                                       `json:"contentId"`
	Content             *content.Content                           `json:"content"`
	Status              UserPracticeStatus                         `json:"status" gorm:"type:user_practice_status"`
	Score               float32                                    `json:"score"`
	Comment             *string                                    `json:"comment"`
	UserPracticeRecords *[]user_practice_record.UserPracticeRecord `json:"userPracticeRecords"`
}

type UserPracticeStatus string

const (
	InProgress UserPracticeStatus = "in progress"
	Submitted  UserPracticeStatus = "submitted"
	Reviewed   UserPracticeStatus = "reviewed"
)
