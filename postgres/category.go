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
		return model.CategoryItem{}, sql.ErrNoRows
	}

	if err != nil {
		return model.CategoryItem{}, err
	}

	return category, nil
}
