package seeder

import (
	"context"

	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/lib"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role"
	roleRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/role/repository"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user"
	userRepository "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user/repository"
)

type seeder struct {
	userRepo userRepository.Repository
	roleRepo roleRepository.Repository
}

func RunSeeders(
	userRepo userRepository.Repository,
	roleRepo roleRepository.Repository,
) {
	s := &seeder{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}

	s.runRoleSeeder()
	s.runUserSeeder()
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
