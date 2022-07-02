package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"library/commons/types"
	"library/model"

	"github.com/labstack/echo/v4"
)

func (app *Application) SaveBook(c echo.Context) error {
	var (
		err          error
		author       model.AuthorItem
		countryItem  model.CountryItem
		categoryItem model.CategoryItem
	)

	// bind request
	u := new(types.Book)
	if err := c.Bind(u); err != nil {
		c.Echo().Logger.Error("error binding book request:", err)
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}
	}

	// validator
	if err = c.Validate(u); err != nil {
		c.Echo().Logger.Error("error validating request body:", err)
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "error validating request body",
			Internal: err,
		}
	}

	// check if author is configured at our end or not
	author, err = app.models.Author.GetAuthorByName(u.Author)
	if err != nil {
		c.Echo().Logger.Error("from get author by name:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "author not configured at our end",
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}

	}

	// check if country is configured at our end or not
	countryItem, err = app.models.Country.GetCountryByName(u.Country)
	if err != nil {
		c.Echo().Logger.Error("error from get country by name:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "country not configured at our end",
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}

	}

	// check if category is configured at our end or not
	categoryItem, err = app.models.Category.GetCategoryByName(u.Category)
	if err != nil {
		c.Echo().Logger.Error("error from get category by name:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "category not configured at our end",
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}

	}

	book := model.BookItem{
		Name:     u.Name,
		Author:   author.ID,
		Country:  countryItem.ID,
		Category: categoryItem.ID,
		Price:    u.Price,
	}

	// if both configured save book details
	_, err = app.models.Book.SaveBook(book)
	if err != nil {
		c.Echo().Logger.Error("error saving book:", err)
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}
	}

	c.Echo().Logger.Info("book details saved")
	// Return Success message
	return c.JSON(http.StatusOK, "book saved successfully!")
}

func (app *Application) GetBook(c echo.Context) error {
	var (
		err         error
		bookDetails model.BookDetails
	)

	// Boook name from path param from book/:name
	bookName := c.Param("name")

	// get book details by name
	bookDetails, err = app.models.Book.GetBookDetailsByName(bookName)
	if err != nil {
		c.Echo().Logger.Errorf("error getting book with name:%s err:%v", bookName, err)
		if errors.Is(err, sql.ErrNoRows) {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "currently book is not available",
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}
	}

	c.Echo().Logger.Info("get book details successful")

	// Return Success message
	return c.JSON(http.StatusOK, bookDetails)
}

func (app *Application) ShowBook(c echo.Context) error {
	var (
		err   error
		books []model.BookDetails
	)

	// Pattern from query param book?name=dummy
	pattern := c.QueryParam("name")

	// Check if pattern matches any book name, give all that matches
	books, err = app.models.Book.GetAllBooksByPattern(pattern)
	if err != nil {
		c.Echo().Logger.Errorf("error getting book with pattern:%s err:%v", pattern, err)
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}
	}

	if books == nil {
		c.Echo().Logger.Errorf("no books available that matches the pattern:%s", pattern)
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "no book availabe that matches given pattern",
		}
	}

	c.Echo().Logger.Info("get book details successful")

	// Return Success message
	return c.JSON(http.StatusOK, books)
}
