package content

import (
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/basic_model"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/question"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_practice"
	"github.com/HSI-Sandbox-Golang-Team-2025/Team-2/pkg/user_project"
)

type Content struct {
	basic_model.BasicModel
	TrackID       uint                          `json:"trackId" example:"1" extensions:"x-order=04"`
	Title         string                        `json:"title" example:"Latihan 2" extensions:"x-order=05"`
	Body          string                        `json:"body" example:"Latihan ini anda akan menguji kemampuan React anda" extensions:"x-order=06"`
	Type          ContentType                   `json:"type" gorm:"type:content_types" example:"practice" extensions:"x-order=07"`
	Order         uint                          `json:"order" example:"1" extensions:"x-order=08"`
	Questions     *[]question.Question          `json:"questions" extensions:"x-order=09"`
	IsCompleted   bool                          `json:"isCompleted" gorm:"<-:false;-:migration" example:"false" extensions:"x-order=10"`
	UserPractices *[]user_practice.UserPractice `json:"userPractices" extensions:"x-order=11"`
	UserProjects  *[]user_project.UserProject   `json:"userProjects" etensions:"x-order=12"`
}

type ContentType string

const (
	ContentTypeMaterial ContentType = "material"
	ContentTypePractice ContentType = "practice"
	ContentTypeProject  ContentType = "project"
)
