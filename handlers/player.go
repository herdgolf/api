package handlers

import (
	"net/http"

	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/models"
	"github.com/labstack/echo/v4"
)

func CreatePlayerHandler(c echo.Context) error {
	p := new(models.Player)
	db := db.DB()

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	player := &models.Player{
		Name:    p.Name,
		Profile: models.Profile{Link: "None"},
		Rounds:  []models.Round{},
	}

	if err := db.Create(&player).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &player,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdatePlayerHandler(c echo.Context) error {
	id := c.Param("id")
	p := new(models.Player)
	db := db.DB()

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	existingPlayer := new(models.Player)

	if err := db.First(&existingPlayer, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	existingPlayer.Name = p.Name
	if err := db.Save(&existingPlayer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &existingPlayer,
	}

	return c.JSON(http.StatusOK, response)
}

func GetPlayerHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	var players []*models.Player

	if res := db.Find(&players, id); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": players,
	}

	return c.JSON(http.StatusOK, response)
}

func DeletePlayerHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	player := new(models.Player)

	err := db.Delete(&player, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"message": "Player deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}
