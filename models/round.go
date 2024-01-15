package models

import "gorm.io/gorm"

type Round struct {
	gorm.Model
	Name     string   `json:"name"`
	Players  []Player `gorm:"many2many:rounds_players;" json:"players"`
	Course   Course
	CourseID uint `json:"course_id"`
	Finished bool `gorm:"default:false"`
}
