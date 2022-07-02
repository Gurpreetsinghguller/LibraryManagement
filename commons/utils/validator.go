package utils

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		fmt.Println("Hello error", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
