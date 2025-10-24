package role

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/acl"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
)

type Role struct {
	basic_model.BasicModel
	Name string     `json:"name" example:"mentor"`
	ACLs *[]acl.ACL `json:"acls"`
}

type RoleId uint

const (
	Mentor RoleId = 1
	Santri RoleId = 2
)
