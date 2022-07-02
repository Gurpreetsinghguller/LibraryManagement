package router

import (
	"library/controller"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(app *controller.Application, e *echo.Echo) {
	e.GET("/healthcheck", app.HealthCheck)

	e.POST("/book", app.SaveBook)
	e.GET("/book/:name", app.GetBook)
	e.GET("/book", app.ShowBook)

	e.GET("/author", app.ShowAuthor)

	e.GET("/category", app.ShowCategories)
}
