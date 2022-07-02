package postgres

import (
	"database/sql"
	"errors"

	"library/model"
)

type CategoryManager struct {
	app *Application
}

func (cm CategoryManager) GetCategoryByName(name string) (model.CategoryItem, error) {
	var category model.CategoryItem

	query := `SELECT id, name FROM category c WHERE c.name=$1`

	err := cm.app.db.DB.Get(&category, query, name)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return category, sql.ErrNoRows
	}

	if err != nil {
		return category, err
	}

	return category, nil
}

func (cm CategoryManager) GetCategoryDetails() ([]model.CategoryDetails, error) {
	var categories []model.CategoryDetails

	query := `SELECT c.name FROM category c`

	err := cm.app.db.DB.Select(&categories, query)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return categories, sql.ErrNoRows
	}

	if err != nil {
		return categories, err
	}

	return categories, nil
}
