package user_practice

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice_record"
)

type UserPractice struct {
	basic_model.BasicModel
	UserID uint `json:"userId" example:"1" extensions:"x-order=04"`
	// User                *user.User                                 `json:"user" swaggertype:"string" extensions:"x-order=05"`
	ContentID           uint                                       `json:"contentId" example:"1" extensions:"x-order=06"`
	Status              UserPracticeStatus                         `json:"status" example:"opened" gorm:"type:user_practice_status" extensions:"x-order=07"`
	Score               float32                                    `json:"score" example:"100" extensions:"x-order=08"`
	Comment             *string                                    `json:"comment" example:"Good job!" extensions:"x-order=09"`
	UserPracticeRecords *[]user_practice_record.UserPracticeRecord `json:"userPracticeRecords" extensions:"x-order=10"`
}

type UserPracticeStatus string

const (
	// Opened     UserPracticeStatus = "opened"
	InProgress UserPracticeStatus = "in progress"
	Submitted  UserPracticeStatus = "submitted"
	Reviewed   UserPracticeStatus = "reviewed"
)
