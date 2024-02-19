package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/models"
	"github.com/herdgolf/api/views/player"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PlayerService interface {
	GetAllPlayers() ([]models.Player, error)
}

type PlayerHandler struct {
	PlayerService PlayerService
}

func New(ps PlayerService) *PlayerHandler {
	return &PlayerHandler{
		PlayerService: ps,
	}
}

func GetPlayersHandler(c echo.Context) error {
	db := db.DB()

	var players []*models.Player

	if res := db.Find(&players); res.Error != nil {
		return res.Error
	}

	pi := player.ShowIndex("| Home", player.Show(players))
	return View(c, pi)
}

func View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func GetPlayerDetailsHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	var players *models.Player
	if res := db.Find(&players, id); res.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, res.Error)
	}

	di := player.DetailsIndex(
		fmt.Sprintf(
			"| User details: %s",
			cases.Title(language.English).String(players.Name),
		),
		player.Details(tz, players),
	)
	return View(c, di)
}

// func (ph *PlayerHandler) HandlerShowPlayers(c echo.Context) error {
// 	pdata, err := ph.PlayerService.GetAllPlayers()
// 	if err != nil {
// 		return err
// 	}
//
// 	pi := player.ShowIndex("| Home", player.Show(pdata))
// 	return ph.View(c, pi)
// }
//
// func (ph *PlayerHandler) View(c echo.Context, cmp templ.Component) error {
// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
//
// 	return cmp.Render(c.Request().Context(), c.Response().Writer)
// }
