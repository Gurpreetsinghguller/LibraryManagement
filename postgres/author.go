package postgres

import (
	"errors"
	"fmt"

	"library/model"
)

type AuthorManager struct {
	app *Application
}

func (am AuthorManager) GetAuthorByName(name string) (model.AuthorItem, error) {
	var author model.AuthorItem

	query := fmt.Sprintf("SELECT id, name, country ,category FROM authors a WHERE a.name=$%v", name)

	row := am.app.db.DB.QueryRow(query)
	err := row.Scan(author)

	if err != nil && errors.Is(err, errors.New("")) {
		return model.AuthorItem{}, err
	}
	return author, nil
}
