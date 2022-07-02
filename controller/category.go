package controller

import (
	"net/http"

	"library/model"

	"github.com/labstack/echo/v4"
)

func (app *Application) ShowCategories(c echo.Context) error {
	var (
		err        error
		categories []model.CategoryDetails
	)

	// get available categories
	categories, err = app.models.Category.GetCategoryDetails()
	if err != nil {
		c.Echo().Logger.Errorf("error getting available categories err:%v", err)
		return c.JSON(http.StatusInternalServerError, "internal error")
	}

	c.Echo().Logger.Info("get category details successful")

	// Return Success message
	return c.JSON(http.StatusOK, categories)
}
