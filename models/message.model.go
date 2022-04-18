package models

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ResponseValidateError(a error) error {
	report, ok := a.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, a.Error())
	}

	if castedObject, ok := a.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "len":
				report.Message = fmt.Sprintf("%s value length must be %s",
					err.Field(), err.Param())
			case "numeric":
				report.Message = fmt.Sprintf("%s value must be numeric",
					err.Field())
			}
		}
	}

	return report
}
