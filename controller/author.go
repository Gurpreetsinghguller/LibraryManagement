package controller

import (
	"net/http"

	"library/model"

	"github.com/labstack/echo/v4"
)

func (app *Application) ShowAuthor(c echo.Context) error {
	var (
		err     error
		authors []model.AuthorDetails
	)

	// get authors list
	authors, err = app.models.Author.GetAllAuthors()
	if err != nil {
		c.Echo().Logger.Errorf("error getting authors err:%v", err)
		return c.JSON(http.StatusInternalServerError, "internal error")
	}

	c.Echo().Logger.Info("get authors successful")

	// Return Success message
	return c.JSON(http.StatusOK, authors)
}
