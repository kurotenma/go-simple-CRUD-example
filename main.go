package main

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/advanced/route/game"
	gameRouteSimple "golang-ecommerce-example/internal/simple/route/game"
)

func main() {
	e := echo.New()

	gameRoute.RegisterRoutes(e)
	gameRouteSimple.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))
}
