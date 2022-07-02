package postgres

import (
	"database/sql"
	"errors"

	"library/model"
)

type AuthorManager struct {
	app *Application
}

func (am AuthorManager) GetAuthorByName(name string) (model.AuthorItem, error) {
	var author model.AuthorItem

	query := `SELECT id, name, country ,category FROM
	 authors a WHERE a.name = $1`

	err := am.app.db.DB.Get(&author, query, name)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.AuthorItem{}, sql.ErrNoRows
	}

	if err != nil {
		return model.AuthorItem{}, err
	}

	return author, nil
}
