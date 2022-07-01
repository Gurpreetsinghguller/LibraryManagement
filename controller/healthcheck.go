package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *Application) HealthCheck(c echo.Context) error {
	c.Logger().Print("Health Check Api Success")
	return c.String(http.StatusOK, "Health check api success!")
}
