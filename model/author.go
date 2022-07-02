package model

import "github.com/google/uuid"

type AuthorItem struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Country  int       `db:"country"`
	Category int       `db:"category"`
}
type AuthorManager interface {
	GetAuthorByName(name string) (AuthorItem, error)
}
