package handlers

import (
	"github.com/Utility-Gods/photoship-go/template/auth_views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	homeView := auth_views.Home("Hello Photoship")
	return renderView(c, auth_views.HomeIndex(
		"| Home",
		"",
		false,
		false,
		[]string{},
		[]string{},
		homeView,
	))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
