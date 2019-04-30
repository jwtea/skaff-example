package main

import (
	"net/http"

	"github.com/labstack/echo"
)

//HomeHandler handle home requests
func (app *App) HomeHandler(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "home", "")
}
