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

func (cm BookManager) GetBookDetailsByName(name string) (model.BookDetails, error) {
	var bookDetails model.BookDetails

	query := `SELECT b.name, a.name as author_name,b.price,c1.name as country_name,c2.name as category_name
			FROM book b 
			INNER JOIN authors a ON a.id= b.author_id
			INNER JOIN country c1 ON c1.id=b.published_country
			INNER JOIN category c2 ON c2.id=b.category
			WHERE b.name=$1`

	err := cm.app.db.DB.Get(&bookDetails, query, name)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.BookDetails{}, sql.ErrNoRows
	}

	if err != nil {
		return model.BookDetails{}, err
	}

	return bookDetails, nil
}
