package gameRoute

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/advanced/controller/game"
)

func RegisterRoutes(e *echo.Echo) {
	u := gameController.NewController()
	r := e.Group("/advanced/game")
	r.POST("", u.InsertGame)
	r.PUT("/:id", u.UpdateGame)
	r.GET("/:id", u.GetGame)
	r.PUT("/verify/:id", u.VerifyGame)

	r = e.Group("/advanced/games")
	r.GET("", u.GetGames)

	//r.GET("/", )
}
