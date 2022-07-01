package router

import (
	"library/controller"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(app *controller.Application, e *echo.Echo) {
	e.GET("/healthcheck", app.HealthCheck)
}
