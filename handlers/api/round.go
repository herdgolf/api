package api

import (
	"net/http"

	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/models"
	"github.com/labstack/echo/v4"
)

func CreateRoundHandler(c echo.Context) error {
	r := new(models.Round)
	db := db.DB()

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	round := &models.Round{
		CourseID: r.CourseID,
		Players:  r.Players,
	}

	if err := db.Create(&round).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &round,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetRoundHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	var rounds []*models.Round

	// if res := db.Find(&rounds, id); res.Error != nil {
	// 	return c.JSON(http.StatusInternalServerError, res.Error.Error())
	// }

	// if res := db.Model(&models.Round{}).Preload("Players").Preload("Course").Find(&rounds, id); res.Error != nil {
	if res := db.Preload("Players").Preload("Course").Find(&rounds, id); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": rounds,
	}

	return c.JSON(http.StatusOK, response)
}

func GetRoundsHandler(c echo.Context) error {
	database := db.DB()

	var rounds []*models.Round
	database.Scopes(db.Paginate(c)).Find(&rounds)

	return c.JSON(http.StatusOK, rounds)
}
