package router

import (
	"net/http"

	"library/controller"
)

func InitializeRoutes(app *controller.Application) {
	http.Handle("/healthcheck", app.HealthCheck())
}
