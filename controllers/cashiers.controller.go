package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FindAllCashiers(c echo.Context) error {
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	conv_limit, err := strconv.Atoi(limit)
	conv_skip, err := strconv.Atoi(skip)

	result, err := models.FindCashiersAll(conv_limit, conv_skip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCashiers(c echo.Context) error {
	name := c.FormValue("name")
	passcode := c.FormValue("passcode")

	result, err := models.StoreCashiers(name, passcode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCashiers(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	passcode := c.FormValue("passcode")

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
