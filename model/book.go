package model

import "github.com/google/uuid"

type BookItem struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Author   uuid.UUID `db:"author_id"`
	Country  int       `db:"published_country"`
	Category int       `db:"category"`
	Price    float32   `db:"price"`
}
type BookManager interface {
	SaveBook(book BookItem) (BookItem, error)
}
