package user

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_track"
)

type User struct {
	basic_model.BasicModel
	Nip        string                  `json:"nip" gorm:"unique" example:"ARN-2402001"`
	Password   string                  `json:"password"`
	Name       string                  `json:"name" example:"Luthfi"`
	RoleID     uint                    `json:"roleId" example:"1"`
	Role       *role.Role              `json:"role"`
	UserTracks []*user_track.UserTrack `json:"userTracks"`
	// UserContents []*user_content.UserContent `json:"userContents"`
}
