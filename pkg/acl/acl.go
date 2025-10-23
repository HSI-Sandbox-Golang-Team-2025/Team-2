package acl

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
)

type ACL struct {
	basic_model.BasicModel
	RoleID     *uint `json:"roleId" example:"1"`
	EndpointID uint  `json:"endpointId" example:"1"`
}
