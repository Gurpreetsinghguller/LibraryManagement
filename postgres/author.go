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
		return author, sql.ErrNoRows
	}

	if err != nil {
		return model.AuthorItem{}, err
	}

	return author, nil
}

func (am AuthorManager) GetAllAuthors() ([]model.AuthorDetails, error) {
	var authors []model.AuthorDetails

	query := `SELECT  a.name, c1.name as "country" ,c2.name as "category" FROM authors a
	INNER JOIN country c1 ON c1.id=a.country
	INNER JOIN category c2 ON c2.id=a.category
	`

	err := am.app.db.DB.Select(&authors, query)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return authors, sql.ErrNoRows
	}

	if err != nil {
		return authors, err
	}

	return authors, nil
}
