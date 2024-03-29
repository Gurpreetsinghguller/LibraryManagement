package model

type CountryItem struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Currency string `db:"currency"`
}
type CountryManager interface {
	GetCountryByName(name string) (CountryItem, error)
}
