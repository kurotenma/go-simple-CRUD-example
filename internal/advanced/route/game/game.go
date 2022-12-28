package gameRoute

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/advanced/controller/game"
)

func RegisterRoutes(e *echo.Echo) {
	u := gameController.NewController()
	r := e.Group("/advanced/game")
	r.POST("", u.InsertGame)

	r = e.Group("/advanced/games")
	r.GET("", u.GetGames)

	//r.GET("/", )
}
