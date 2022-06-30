package controller

import (
	"booklibrary/commons/utils"
	"booklibrary/postgres"
)

type Application struct {
	db     *utils.Postgres
	models *postgres.Models
}

func New(db *utils.Postgres, models *postgres.Models) *Application {
	app := &Application{
		db:     db,
		models: models,
	}
	return app
}
