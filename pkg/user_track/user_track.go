package user_track

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/track"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_material"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
)

type UserTrack struct {
	basic_model.BasicModel
	UserID uint `json:"userId"`
	// User    *user.User   `json:"user"`
	TrackID       uint                          `json:"trackId"`
	Track         *track.Track                  `json:"track"`
	Status        UserTrackStatus               `json:"status"`
	AverageScore  float32                       `json:"averageScore"`
	UserMaterials *[]user_material.UserMaterial `json:"userMaterials"`
	UserPractices *[]user_practice.UserPractice `json:"userPractice"`
	UserProjects  *[]user_project.UserProject   `json:"userProjects"`
}

type UserTrackStatus string

const (
	InProgress UserTrackStatus = "in progress"
	Completed  UserTrackStatus = "completed"
)
