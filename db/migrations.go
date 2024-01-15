package db

import (
	"github.com/herdgolf/api/models"
)

func AutoMigrate() {
	// database.AutoMigrate(&models.Profile{})
	database.AutoMigrate(&models.Player{})
	database.AutoMigrate(&models.Hole{})
	database.AutoMigrate(&models.Round{})
	database.AutoMigrate(&models.ScoreCard{})
	database.AutoMigrate(&models.HoleScore{})
	// database.AutoMigrate(&models.Course{})
}
