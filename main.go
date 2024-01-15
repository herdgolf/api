package main

import (
	"net/http"

	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/handlers"
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
	userRoute := e.Group("/player")
	userRoute.POST("/", handlers.CreatePlayerHandler)
	userRoute.GET("/:id", handlers.GetPlayerHandler)
	userRoute.PUT("/:id", handlers.UpdatePlayerHandler)
	userRoute.DELETE("/:id", handlers.DeletePlayerHandler)

	scoreCardRoute := e.Group("/scorecard")
	scoreCardRoute.POST("/", handlers.CreateScoreCardHandler)
	scoreCardRoute.GET("/:id", handlers.GetScoreCardHandler)
	scoreCardRoute.GET("/:id/holes", handlers.GetHolesByScoreCardID)
	scoreCardRoute.GET("/:score_card_id/holes/:number", handlers.GetHoleByNumberAndScoreCardID)
	scoreCardRoute.PUT("/:id", handlers.UpdateScoreCardHandler)
	scoreCardRoute.PUT("/score", handlers.AddScoreToScoreCardHandler)

	courseRoute := e.Group("/course")
	courseRoute.POST("/", handlers.CreateCourseHandler)
	courseRoute.GET("/:id", handlers.GetCourseHandler)
	courseRoute.GET("/:id/holes", handlers.GetHolesByCourseIDHandler)
	courseRoute.GET("/:id/holes/:number", handlers.GetHoleByNumberHandler)
	courseRoute.PUT("/:id/holes/:number", handlers.UpdateHoleHandler)
	courseRoute.PUT("/:id", handlers.UpdateCourseHandler)
	courseRoute.DELETE("/:id", handlers.DeleteCourseHandler)

	roundRoute := e.Group("/round")
	roundRoute.POST("/", handlers.CreateRoundHandler)
	roundRoute.GET("/:id", handlers.GetRoundHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
