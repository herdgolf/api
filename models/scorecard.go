package models

import "gorm.io/gorm"

type ScoreCard struct {
	gorm.Model
	Name       string `json:"name"`
	Course     Course
	CourseID   uint `json:"course_id"`
	Player     Player
	PlayerID   uint `json:"player_id"`
	Round      Round
	RoundID    uint        `json:"round_id"`
	HoleScores []HoleScore `json:"hole_scores"`
}

type HoleScore struct {
	HoleID      int  `json:"hole_id"`
	Strokes     int  `json:"strokes"`
	ScoreCardID uint `json:"score_card_id"`
}
