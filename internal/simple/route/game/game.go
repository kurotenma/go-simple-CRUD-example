package gameRouteSimple

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/simple/controller/game"
)

func RegisterRoutes(e *echo.Echo) {
	r := e.Group("/simple/game")
	r.POST("", gameControllerSimple.InsertGame)
	r.GET("/:id", gameControllerSimple.GetGame)
	r.PUT("/:id", gameControllerSimple.UpdateGame)
	r.PUT("/verify/:id", gameControllerSimple.VerifyGame)

	r = e.Group("/simple/games")
	r.GET("", gameControllerSimple.GetGames)
}
