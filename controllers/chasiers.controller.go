package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FindAllChasiers(c echo.Context) error {
	result, err := models.FindChasiersAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreChasiers(c echo.Context) error {
	name := c.FormValue("name")
	passcode := c.FormValue("passcode")

	result, err := models.StoreChasiers(name, passcode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateChasiers(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	passcode := c.FormValue("passcode")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateChasiers(conv_id, name, passcode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteChasiers(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteChasiers(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
