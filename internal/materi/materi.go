package user

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/internal/role"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nip      string    `json:"nip" gorm:"unique"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	RoleID   uint      `json:"roleId"`
	Role     role.Role `json:"role"`
}
