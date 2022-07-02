package router

import (
	"library/controller"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(app *controller.Application, e *echo.Echo) {
	e.GET("/healthcheck", app.HealthCheck)

	// Save book
	e.POST("/book", app.SaveBook)
	// Get specific book /book/:Dummy
	e.GET("/book/:name", app.GetBook)
	// search for books /book?name=Dummy
	e.GET("/book", app.ShowBook)

	// Get all available author list
	e.GET("/author", app.ShowAuthor)
	// Get all category of books
	e.GET("/category", app.ShowCategories)
}
