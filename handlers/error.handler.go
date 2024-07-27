package handlers

import (
	"fmt"

	"github.com/Utility-Gods/photoship-go/template/errors_pages"
	"github.com/labstack/echo/v4"
)

func RouteNotFoundHandler(c echo.Context) error {
	// Hardcoded parameters

	return renderView(c, errors_pages.ErrorIndex(
		fmt.Sprintf("| Error (%d)", 404),
		"",
		false,
		true,
		errors_pages.Error404(false),
	))
}
