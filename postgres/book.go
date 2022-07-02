package postgres

import (
	"database/sql"
	"errors"

	"library/model"
)

type BookManager struct {
	app *Application
}

func (cm BookManager) SaveBook(book model.BookItem) (model.BookItem, error) {
	var category model.BookItem

	query := `INSERT INTO book (name,author_id,category,published_country,price)
	 values ($1,$2,$3,$4,$5)`

	_, err := cm.app.db.DB.Exec(query, book.Name, book.Author,
		book.Category, book.Country, book.Price)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.BookItem{}, sql.ErrNoRows
	}

	if err != nil {
		return model.BookItem{}, err
	}

	return category, nil
}
