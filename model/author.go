package model

type AuthorItem struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Country  int    `db:"country"`
	Category int    `db:"category"`
}
type AuthorManager interface {
	GetAuthorByName(name string) (AuthorItem, error)
}
