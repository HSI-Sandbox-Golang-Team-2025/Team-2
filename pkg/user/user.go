package user

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
)

type User struct {
	basic_model.BasicModel
	Nip      string     `json:"nip" gorm:"unique"`
	Password string     `json:"password"`
	Name     string     `json:"name"`
	RoleID   uint       `json:"roleId"`
	Role     *role.Role `json:"role"`
}
