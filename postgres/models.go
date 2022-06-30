package postgres

import (
	"booklibrary/commons/utils"
	"booklibrary/model"
)

type Application struct {
	db *utils.Postgres
}

type Models struct {
	Book model.BookHelper
}

func NewModels(db *utils.Postgres) *Models {
	app := &Application{db: db}
	return &Models{
		Book: &BookHelper{app},
	}
}
