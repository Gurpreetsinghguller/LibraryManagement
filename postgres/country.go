package postgres

import (
	"errors"
	"fmt"

	"library/model"
)

type CountryManager struct {
	app *Application
}

func (cm CountryManager) GetCountryByName(name string) (model.CountryItem, error) {
	var country model.CountryItem

	query := fmt.Sprintf("SELECT id, name, currency FROM country c WHERE c.name=$%v", name)

	row := cm.app.db.DB.QueryRow(query)
	err := row.Scan(country)
	if err != nil && errors.Is(err, errors.New("")) {
		return model.CountryItem{}, err
	}

	return country, nil
}
