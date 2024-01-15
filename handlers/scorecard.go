package handlers

import (
	"fmt"
	"net/http"

	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func CreateScoreCardHandler(c echo.Context) error {
	s := new(models.ScoreCard)
	db := db.DB()

	if c.Bind(s) != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	scorecard := models.ScoreCard{
		Name:       s.Name,
		PlayerID:   s.PlayerID,
		RoundID:    s.RoundID,
		CourseID:   s.CourseID,
		HoleScores: []models.HoleScore{},
	}
	if err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&scorecard).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &scorecard,
	}

	return c.JSON(http.StatusCreated, response)
}

func AddScoreToScoreCardHandler(c echo.Context) error {
	score := new(models.HoleScore)
	db := db.DB()
	if err := c.Bind(score); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updatedScore := models.HoleScore{
		HoleID:      score.HoleID,
		Strokes:     score.Strokes,
		ScoreCardID: score.ScoreCardID,
	}

	if err := db.Create(&updatedScore).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &updatedScore,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetScoreCardHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	var sc []*models.ScoreCard

	if res := db.Model(&models.ScoreCard{}).Preload("Holes").Find(&sc, id); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": sc,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateScoreCardHandler(c echo.Context) error {
	id := c.Param("id")
	sc := new(models.ScoreCard)
	db := db.DB()

	if err := c.Bind(sc); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	existingScoreCard := new(models.ScoreCard)

	if err := db.First(&existingScoreCard, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	existingScoreCard.Name = sc.Name

	if err := db.Save(&existingScoreCard).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &existingScoreCard,
	}

	return c.JSON(http.StatusOK, response)
}

func GetHolesByScoreCardID(c echo.Context) error {
	scorecardID := c.Param("id")
	db := db.DB()

	var holes []*models.Hole

	if res := db.Find(&holes, "score_card_id = ?", scorecardID); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}
	fmt.Println(holes[0])

	response := map[string]interface{}{
		"data": holes,
	}

	return c.JSON(http.StatusOK, response)
}

func GetHoleByNumberAndScoreCardID(c echo.Context) error {
	scorecardID := c.Param("score_card_id")
	holeID := c.Param("number")
	db := db.DB()

	var hole *models.Hole

	if res := db.Find(&hole, "score_card_id = ? AND number = ?", scorecardID, holeID); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": hole,
	}

	return c.JSON(http.StatusOK, response)
}
