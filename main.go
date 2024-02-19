package main

import (
	"net/http"

	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/handlers"
	"github.com/herdgolf/api/handlers/api"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	dbGorm.Ping()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	e.Static("/", "assets")
	playerRoute := e.Group("/player")
	playerRoute.GET("/", handlers.GetPlayersHandler)
	playerRoute.GET("/details/:id", handlers.GetPlayerDetailsHandler)

	userRoute := e.Group("/api/player")
	userRoute.POST("/", api.CreatePlayerHandler)
	userRoute.GET("/:id", api.GetPlayerHandler)
	userRoute.GET("/", api.GetPlayersHandler)
	userRoute.PUT("/:id", api.UpdatePlayerHandler)
	userRoute.DELETE("/:id", api.DeletePlayerHandler)

	scoreCardRoute := e.Group("/api/scorecard")
	scoreCardRoute.POST("/", api.CreateScoreCardHandler)
	scoreCardRoute.GET("/", api.GetScoreCardsHandler)
	scoreCardRoute.GET("/:id", api.GetScoreCardHandler)
	scoreCardRoute.GET("/:id/holes", api.GetHolesByScoreCardID)
	scoreCardRoute.GET("/:score_card_id/holes/:number", api.GetHoleByNumberAndScoreCardID)
	scoreCardRoute.PUT("/:id", api.UpdateScoreCardHandler)
	scoreCardRoute.PUT("/score", api.AddScoreToScoreCardHandler)

	courseRoute := e.Group("/api/course")
	courseRoute.POST("/", api.CreateCourseHandler)
	courseRoute.GET("/:id", api.GetCourseHandler)
	courseRoute.GET("/", api.GetCoursesHandler)
	courseRoute.GET("/:id/holes", api.GetHolesByCourseIDHandler)
	courseRoute.GET("/:id/holes/:number", api.GetHoleByNumberHandler)
	courseRoute.PUT("/:id/holes/:number", api.UpdateHoleHandler)
	courseRoute.PUT("/:id", api.UpdateCourseHandler)
	courseRoute.DELETE("/:id", api.DeleteCourseHandler)

	roundRoute := e.Group("/api/round")
	roundRoute.POST("/", api.CreateRoundHandler)
	roundRoute.GET("/:id", api.GetRoundHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
