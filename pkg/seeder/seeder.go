package seeder

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/lib"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/acl"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint"
	endpointRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/endpoint/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	roleRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	trackRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	userRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/repository"
)

type seeder struct {
	userRepo     userRepository.Repository
	roleRepo     roleRepository.Repository
	endpointRepo endpointRepository.Repository
	trackRepo    trackRepository.Repository
}

func RunSeeders(
	userRepo userRepository.Repository,
	roleRepo roleRepository.Repository,
	endpointRepo endpointRepository.Repository,
	trackRepo trackRepository.Repository,
) {
	s := &seeder{
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		endpointRepo: endpointRepo,
		trackRepo:    trackRepo,
	}

	s.runRoleSeeder()
	s.runUserSeeder()
	s.runEndpointSeeder()
	s.runTrackSeeder()
}

func (s *seeder) runRoleSeeder() error {
	roles := []role.Role{
		{
			BasicModel: basic_model.BasicModel{
				ID: 1,
			},
			Name: "mentor",
		},
		{
			BasicModel: basic_model.BasicModel{
				ID: 2,
			},
			Name: "santri",
		},
	}

	for _, r := range roles {
		s.roleRepo.CreateRole(context.Background(), &r)
	}

	return nil
}

func (s *seeder) runUserSeeder() error {
	users := []user.User{
		{
			Nip:      "ARN-2402001",
			RoleID:   1,
			Name:     "Luthfi",
			Password: "123",
		},
	}

	for _, u := range users {
		hassPassword, err := lib.HashPassword(u.Password)

		if err != nil {
			return err
		}

		u.Password = hassPassword

		s.userRepo.CreateUser(context.Background(), &u)
	}

	return nil
}

func (s *seeder) runEndpointSeeder() error {
	endpoints := []endpoint.Endpoint{
		{
			Path:   "/api/auth/register",
			Method: "POST",
		},
		{
			Path:   "/api/auth/login",
			Method: "POST",
		},
		{
			Path:   "/api/contents",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/contents",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/contents/:id",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/contents/:id",
			Method: "PUT",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/contents/:id",
			Method: "DELETE",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/materials",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/materials",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/materials/:id",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/materials/:id",
			Method: "PUT",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/materials/:id",
			Method: "DELETE",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/practices",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/projects",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/tracks",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/tracks",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/tracks/:id",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/tracks/:id",
			Method: "PUT",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/tracks/:id",
			Method: "DELETE",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/user-materials",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/user-materials",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-materials/:id",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-materials/:id",
			Method: "PUT",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/user-materials/:id",
			Method: "DELETE",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
			},
		},
		{
			Path:   "/api/user-practices",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-practices/:id/start",
			Method: "PATCH",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-practices/:id/submit",
			Method: "PATCH",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-practices/:id/review",
			Method: "PATCH",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-projects",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-projects",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-projects/:id/submit",
			Method: "PATCH",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-projects/:id/review",
			Method: "PATCH",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-tracks",
			Method: "POST",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-tracks",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-tracks/:id",
			Method: "GET",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-tracks/:id",
			Method: "PUT",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
		{
			Path:   "/api/user-tracks/:id",
			Method: "DELETE",
			ACLs: &[]acl.ACL{
				{RoleID: 1},
				{RoleID: 2},
			},
		},
	}

	for _, e := range endpoints {
		s.endpointRepo.CreateEndpoint(context.Background(), &e)
	}

	return nil
}

func (s *seeder) runTrackSeeder() error {
	tracks := []track.Track{
		{Name: "Flutter Development"},
		{Name: "NextJs Development"},
		{Name: "Golang Development"},
		{Name: "Phyton (Django) Development"},
		{Name: "UI/UX Design"},
	}

	for _, t := range tracks {
		s.trackRepo.CreateTrack(context.Background(), &t)
	}

	return nil
}
