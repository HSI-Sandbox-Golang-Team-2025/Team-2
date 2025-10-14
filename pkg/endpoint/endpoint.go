package endpoint

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/acl"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
)

type Endpoint struct {
	basic_model.BasicModel
	Method string     `json:"method"`
	Path   string     `json:"path"`
	ACLs   *[]acl.ACL `json:"acls"`
}
