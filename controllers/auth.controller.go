package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func VerifyLogin(c echo.Context) error {
	id := c.Param("id")
	cashier := &models.CashiersValidation{}

	err := c.Bind(cashier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

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

func VerifyLogout(c echo.Context) error {
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
