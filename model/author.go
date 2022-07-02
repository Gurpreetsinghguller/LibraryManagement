package model

import "github.com/google/uuid"

type AuthorItem struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Country  int       `db:"country"`
	Category int       `db:"category"`
}

type AuthorDetails struct {
	Name     string `db:"name"`
	Country  string `db:"country"`
	Category string `db:"category"`
}
type AuthorManager interface {
	GetAuthorByName(name string) (AuthorItem, error)
	GetAllAuthors() ([]AuthorDetails, error)
}
