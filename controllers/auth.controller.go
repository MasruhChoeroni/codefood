package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginCashiers(c echo.Context) error {
	id := c.Param("id")
	cashier := &models.Cashiers{}

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.Bind(cashier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	httpNumber, result, err := models.VerifyLogin(conv_id, cashier.Passcode)

	if err != nil {
		return c.JSON(httpNumber.Number, map[string]string{"message": err.Error()})
	}

	return c.JSON(httpNumber.Number, result)
}
