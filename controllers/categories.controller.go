package controllers

import (
	"codefood/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FindAllCategories(c echo.Context) error {
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

	result, err := models.FindCategoriesAll(conv_limit, conv_skip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindCategoriesById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCategoriesById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCategories(c echo.Context) error {
	name := c.FormValue("name")

	result, err := models.StoreCategories(name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCategories(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateCategories(conv_id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCategories(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id) //convert to integer
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteCategories(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
