package controllers

import (
	"codefood/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func FindAllCashiers(c echo.Context) error {
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	conv_limit, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	conv_skip, err := strconv.Atoi(skip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCashiersAll(conv_limit, conv_skip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindCashiersById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCashiersById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindCashiersPasscodeById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCashiersPasscodeById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCashiers(c echo.Context) error {
	validate := validator.New()
	cashier := &models.CashiersPostValidation{}

	err := c.Bind(cashier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(cashier)

	if err != nil {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
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
		return c.JSON(http.StatusInternalServerError, report)
	}

	result, err := models.StoreCashiers(cashier.Name, cashier.Passcode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCashiers(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	passcode := c.FormValue("passcode")

	validate = validator.New()

	type CashiersValidate struct {
		Name     string `validate:"required"`
		Passcode string `validate:"required,numeric,len=6"`
	}

	outer := &CashiersValidate{
		Name:     name,
		Passcode: passcode,
	}

	err := validate.Struct(outer)

	if err != nil {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
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
		return c.JSON(http.StatusInternalServerError, report)
	}

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateCashiers(conv_id, name, passcode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCashiers(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteCashiers(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
