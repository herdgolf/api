package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name  string `json:"name"`
	Par   int    `json:"par"`
	Holes []Hole `json:"holes"`
}

type Hole struct {
	gorm.Model
	Number      int `json:"number"`
	Par         int `json:"par"`
	StrokeIndex int `json:"stroke_index"`
	CourseID    uint
}
