package services

import (
	"time"

	"github.com/herdgolf/api/models"
	"gorm.io/gorm"
)

func NewServicesPlayer(p models.Player, db *gorm.DB) *ServicesPlayer {
	return &ServicesPlayer{p, db}
}

type ServicesPlayer struct {
	Player models.Player
	db     *gorm.DB
}

func (sp *ServicesPlayer) GetAllPlayers() ([]*models.Player, error) {
	var players []*models.Player

	if res := sp.db.Find(&players); res.Error != nil {
		return nil, res.Error
	}

	return players, nil
}

func ConverDateTime(tz string, dt time.Time) string {
	loc, _ := time.LoadLocation(tz)

	return dt.In(loc).Format(time.RFC822Z)
}
