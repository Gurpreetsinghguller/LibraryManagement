package postgres

import (
	"database/sql"
	"errors"

	"library/model"
)

type CountryManager struct {
	app *Application
}

func (cm CountryManager) GetCountryByName(name string) (model.CountryItem, error) {
	var country model.CountryItem

	query := `SELECT id, name, currency FROM country c WHERE c.name=$1`

	err := cm.app.db.DB.Get(&country, query, name)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.CountryItem{}, sql.ErrNoRows
	}

	if err != nil {
		return model.CountryItem{}, err
	}

	return country, nil
}
