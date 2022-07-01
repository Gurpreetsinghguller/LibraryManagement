package model

type BookItems struct {
	ID       string  `db:"id"`
	Name     string  `db:"name"`
	Author   string  `db:"author"`
	Country  int     `db:"country"`
	Category int     `db:"category"`
	Price    float32 `db:"price"`
}
type BookManager interface {
	// Get Books
	// Save Books
}
