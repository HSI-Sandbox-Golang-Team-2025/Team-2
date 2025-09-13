package track

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name"`
}

func (Track) TableName() string {
	return "tracks"
}
