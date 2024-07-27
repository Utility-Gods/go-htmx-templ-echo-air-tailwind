package main

import (
	"github.com/Utility-Gods/photoship-go/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// component := template.Home("PhotoShip")

	e := echo.New()
	e.Use(middleware.Logger())


	handlers.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
