package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name      string `json:"name"`
	ProfileID int
	Profile   Profile
	Rounds    []Round `gorm:"many2many:rounds_players;" json:"rounds"`
}

type Profile struct {
	Link string `json:"link"`
	ID   int
}
