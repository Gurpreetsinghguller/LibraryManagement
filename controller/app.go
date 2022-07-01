package controller

import (
	"library/db/dbconn"
	"library/postgres"
)

type Application struct {
	db     *dbconn.Postgres
	models *postgres.Models
}

func New(db *dbconn.Postgres, models *postgres.Models) *Application {
	app := &Application{
		db:     db,
		models: models,
	}
	return app
}
