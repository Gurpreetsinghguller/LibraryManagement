package controller

import (
	"net/http"
)

func (app *Application) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success Health check api"))
	}
}
