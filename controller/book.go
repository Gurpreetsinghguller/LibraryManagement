package controller

import (
	"net/http"

	"library/commons/types"

	"github.com/labstack/echo/v4"
)

func (app *Application) SaveBook(c echo.Context) error {
	var err error
	u := new(types.Book)
	if err := c.Bind(u); err != nil {
		return err
	}

	// check if country is configured at our end or not
	_, err = app.models.Country.GetCountryByName(u.Country)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Country not configured at our end")
	}
	// check if author is configured at our end or not
	_, err = app.models.Author.GetAuthorByName(u.Author)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Author not configured at our end")
	}

	// if both configured save book details

	// Return Success message
	return c.JSON(http.StatusCreated, u)
}

func (app *Application) GetBook(c echo.Context) error {
	c.Logger().Print("Health Check Api Success")
	return c.String(http.StatusOK, "Health check api success!")
}
