package model

type CategoryItem struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type CategoryDetails struct {
	Name string `db:"name"`
}
type CategoryManager interface {
	GetCategoryByName(name string) (CategoryItem, error)
	GetCategoryDetails() ([]CategoryDetails, error)
}
