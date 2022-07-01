package postgres

import (
	"library/commons/dbconn"
	"library/model"
)

type Application struct {
	db *dbconn.Postgres
}

type Models struct {
	Book model.BookHelper
}

func NewModels(db *dbconn.Postgres) *Models {
	app := &Application{db: db}
	return &Models{
		Book: &BookHelper{app},
	}
}
