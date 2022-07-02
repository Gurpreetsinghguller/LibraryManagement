package postgres

import (
	"library/db/dbconn"
	"library/model"
)

type Application struct {
	db *dbconn.Postgres
}

type Models struct {
	Book     model.BookManager
	Country  model.CountryManager
	Author   model.AuthorManager
	Category model.CategoryManager
}

func NewModels(db *dbconn.Postgres) *Models {
	app := &Application{db: db}
	return &Models{
		Book:     &BookManager{app},
		Country:  &CountryManager{app},
		Author:   &AuthorManager{app},
		Category: &CategoryManager{app},
	}
}
